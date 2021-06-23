package modules

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/icetrays/icetrays/consensus"
	"github.com/icetrays/icetrays/consensus/pb"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Op struct {
	Op     string   `json:"op"`
	Params []string `json:"params"`
}

func Server2(node *consensus.Node, config Config) error {
	router := gin.Default()
	mulAddr, err := multiaddr.NewMultiaddr(config.Ipfs)
	if err != nil {
		return err
	}
	netAddr, err := manet.ToNetAddr(mulAddr)
	if err != nil {
		return err
	}
	ipfsUrl, err := url.Parse(fmt.Sprintf("http://%s", netAddr.String()))
	if err != nil {
		return err
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(ipfsUrl)
	reverseProxy.Transport = http.DefaultTransport

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.POST("/icetrays", func(c *gin.Context) {
		d, err := c.GetRawData()
		if err != nil {
			return
		}
		op := &Op{}
		err = json.Unmarshal(d, op)
		if err != nil {
			c.JSON(200, err.Error())
		}
		switch op.Op {
		case "ls":
			n, _ := node.Ls(c, op.Params[0])
			c.JSON(200, n)
		case "cp":
			err := node.Op(c, pb.Instruction_CP, op.Params[0], op.Params[1])
			if err != nil {
				c.JSON(200, err.Error())
			} else {
				c.JSON(200, "success")
			}
		case "mv":
			err := node.Op(c, pb.Instruction_MV, op.Params[0], op.Params[1])
			if err != nil {
				c.JSON(200, err.Error())
			} else {
				c.JSON(200, "success")
			}
		case "rm":
			err := node.Op(c, pb.Instruction_RM, op.Params[0])
			if err != nil {
				c.JSON(200, err.Error())
			} else {
				c.JSON(200, "success")
			}
		case "mkdir":
			err := node.Op(c, pb.Instruction_MKDIR, op.Params[0])
			if err != nil {
				c.JSON(200, err.Error())
			} else {
				c.JSON(200, "success")
			}
		default:
			c.JSON(200, "???")
		}
	})
	var proxyHandle = func(c *gin.Context) {
		reverseProxy.ServeHTTP(c.Writer, c.Request)
	}
	router.POST("/", proxyHandle)
	router.GET("/", proxyHandle)
	go router.Run(fmt.Sprintf(":%d", config.Port))
	return nil
}

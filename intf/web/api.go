package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/bacon-number-go/entity"
)

type Application struct {
	nodeStore  entity.NodeStore
	tableStore entity.TableStore
	// queue      entity.Queue
	dvr *entity.DVRImpl
}

func getNodes(app *Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		nodes := []*entity.Node{}
		if err := app.nodeStore.GetNodes(&nodes); err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		ret := []*Node{}
		for _, n := range nodes {
			tbl := entity.Table{}
			if err := app.tableStore.GetTable(ctx, n.ID, &tbl); err != nil {
				if err != entity.NotExistErr {
					ctx.AbortWithError(500, err)
					return
				}
			}
			ret = append(ret, NewNode(n, &tbl))
		}
		ctx.IndentedJSON(200, ret)
		return
	}
}

func postNodes(app *Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id := entity.NodeID(ctx.Param("id"))
		node := entity.Node{}
		if err := app.nodeStore.GetNode(id, &node); err != nil {
			if err != entity.NotExistErr {
				ctx.AbortWithError(500, err)
				return
			}
		} else {
			ctx.AbortWithError(500, fmt.Errorf("Already exist node"))
			return
		}
		if err := app.nodeStore.PutNode(entity.NewNode(id)); err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		ctx.Status(200)
		return
	}
}

func putEdges(app *Application) func(*gin.Context) {
	return func(ctx *gin.Context) {
		tailID := entity.NodeID(ctx.Param("tailID"))
		headID := entity.NodeID(ctx.Param("headID"))
		if err := app.nodeStore.PutEdge(tailID, headID, true); err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		if err := app.dvr.Next(ctx, tailID); err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		ctx.Status(200)
		return
	}
}

func NewRoute(
	nodeStore entity.NodeStore,
	tableStore entity.TableStore,
	dvr *entity.DVRImpl,
) *gin.Engine {
	app := Application{
		nodeStore:  nodeStore,
		tableStore: tableStore,
		dvr:        dvr,
	}
	r := gin.Default()
	r.GET("/nodes", getNodes(&app))
	r.POST("/nodes/:id", postNodes(&app))
	r.PUT("/edges/:tailID/:headID", putEdges(&app))
	return r
}

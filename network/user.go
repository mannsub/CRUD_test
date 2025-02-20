package network

import (
	"sync"

	"github.com/Mansub/CRUD_SERVER/service"
	"github.com/Mansub/CRUD_SERVER/types"
	"github.com/gin-gonic/gin"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network

	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		router.registerPOST("/", userRouterInstance.create)
		router.registerGET("/", userRouterInstance.get)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)
	})

	return userRouterInstance
}

func (u *userRouter) create(c *gin.Context) {
	var req types.CreateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			APIResponse: types.NewAPIResponse("Binding error.", -1, err.Error()),
		})
	} else if err = u.userService.Create(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			APIResponse: types.NewAPIResponse("Create error.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateUserResponse{
			APIResponse: types.NewAPIResponse("Success.", 1, nil),
		})
	}
}

func (u *userRouter) get(c *gin.Context) {
	u.router.okResponse(c, &types.GetUserResponse{
		APIResponse: types.NewAPIResponse("Success.", 1, nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			APIResponse: types.NewAPIResponse("Binding error.", -1, err.Error()),
		})
	} else if err = u.userService.Update(req.Name, req.UpdatedAge); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			APIResponse: types.NewAPIResponse("Update error.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.UpdateUserResponse{
			APIResponse: types.NewAPIResponse("Success.", 1, nil),
		})
	}
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			APIResponse: types.NewAPIResponse("Binding error.", -1, err.Error()),
		})
	} else if err = u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.CreateUserResponse{
			APIResponse: types.NewAPIResponse("Delete error.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			APIResponse: types.NewAPIResponse("Success.", 1, nil),
		})
	}
}

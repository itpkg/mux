package mux

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/itpkg/cache"
)


func Cache(store cache.Store, expire uint, handler Handler) Handler{
	return func(ctx *Context) error{
		var cr Response
		key := ctx.Request.URL.Path

		if err := store.Get(key, &cr); err == nil{
			ctx.Write(&cr)
			return nil
		}
		if err := handler(ctx); err == nil{
			store.Set
		}else{
			return err
		}


	}
}

func Token(fn jwt.KeyFunc, handler Handler) Handler{
	return func(ctx *Context) error{
		if tk, err := jwt.ParseFromRequest(ctx.Request, fn); err==nil{
			ctx.Params["token"] = tk
			return handler(ctx)
		}else{
			return err
		}
	}
}

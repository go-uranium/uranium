package uranium

//func (ushio *Ushio) HandlePost(ctx *fiber.Ctx) error {
//	ushio.Lock.RLock()
//	defer ushio.Lock.RUnlock()
//	nav, err := ushio.NavFromCtx(ctx)
//	if err != nil {
//		return err
//	}
//
//	postID := ctx.Params("post")
//	if len(postID) < 1 || len(postID) > 10 {
//		return fiber.NewError(http.StatusBadRequest, "Invalid post id.")
//	}
//
//	pid, err := strconv.Atoi(postID)
//	if err != nil {
//		return fiber.NewError(http.StatusBadRequest, "Invalid post id.")
//	}
//
//	post, err := ushio.Data.PostByPID(int64(pid))
//	if err != nil {
//		if err == sql.ErrNoRows {
//			return fiber.NewError(http.StatusNotFound, "Post not found.")
//		}
//		return err
//	}
//
//	err = ushio.Data.PostNewView(post.PID)
//	if err != nil {
//		return err
//	}
//
//	err = ushio.Cache.IndexPostInfoRefresh()
//	if err != nil {
//		return err
//	}
//
//	comments, err := ushio.Data.CommentsByPost(post.PID)
//	if err != nil && err != sql.ErrNoRows {
//		return err
//	}
//
//	return ctx.Render("post", fiber.Map{
//		"Meta": &Meta{
//			Config:      *ushio.Config,
//			CurrentPage: post.Info.Title,
//		},
//		"Nav":      nav,
//		"Post":     post,
//		"Comments": comments,
//	})
//}

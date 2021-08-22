package uranium

//type VoteType uint8
//
//const (
//	POSITIVE VoteType = iota
//	NEGATIVE
//)
//
//type VoteReq struct {
//	Type VoteType `json:"type"`
//	ID   int64    `json:"id"`
//}
//
//type VoteResp struct {
//	Success bool   `json:"success"`
//	Err     string `json:"err"`
//	Pos     int    `json:"pos"`
//	Neg     int    `json:"neg"`
//}
//
//func (ushio *Ushio) HandlePOSTVotePost(ctx *fiber.Ctx) error {
//	nav, err := ushio.NavFromCtx(ctx)
//	if err != nil {
//		return err
//	}
//	if !nav.LoggedIn {
//		return ctx.JSON(&VoteResp{
//			Success: false,
//			Err:     "Not logged in.",
//		})
//	}
//	req := &VoteReq{}
//	err = ctx.BodyParser(req)
//	if err != nil {
//		return err
//	}
//
//	p, err := ushio.Data.PostInfoByPID(req.ID)
//	if err != nil {
//		return err
//	}
//
//	switch req.Type {
//	case POSITIVE:
//		if p.VoteNegContains(nav.User.UID) {
//			err := ushio.Data.PostRemoveNegVote(p.PID, nav.User.UID)
//			if err != nil {
//				return err
//			}
//			err = ushio.Data.PostNewPosVote(p.PID, nav.User.UID)
//			if err != nil {
//				return err
//			}
//			return ctx.JSON(&VoteResp{
//				Success: true,
//				Pos:     p.VotePosCount() + 1,
//				Neg:     p.VoteNegCount() - 1,
//			})
//		}
//
//		if p.VotePosContains(nav.User.UID) {
//			err := ushio.Data.PostRemovePosVote(p.PID, nav.User.UID)
//			if err != nil {
//				return err
//			}
//			return ctx.JSON(&VoteResp{
//				Success: true,
//				Pos:     p.VotePosCount() - 1,
//				Neg:     p.VoteNegCount(),
//			})
//		}
//
//		err := ushio.Data.PostNewPosVote(p.PID, nav.User.UID)
//		if err != nil {
//			return err
//		}
//
//		return ctx.JSON(&VoteResp{
//			Success: true,
//			Pos:     p.VotePosCount() + 1,
//			Neg:     p.VoteNegCount(),
//		})
//
//	case NEGATIVE:
//		if p.VotePosContains(nav.User.UID) {
//			err := ushio.Data.PostRemovePosVote(p.PID, nav.User.UID)
//			if err != nil {
//				return err
//			}
//			err = ushio.Data.PostNewNegVote(p.PID, nav.User.UID)
//			if err != nil {
//				return err
//			}
//			return ctx.JSON(&VoteResp{
//				Success: true,
//				Pos:     p.VotePosCount() - 1,
//				Neg:     p.VoteNegCount() + 1,
//			})
//		}
//
//		if p.VoteNegContains(nav.User.UID) {
//			err := ushio.Data.PostRemoveNegVote(p.PID, nav.User.UID)
//			if err != nil {
//				return err
//			}
//			return ctx.JSON(&VoteResp{
//				Success: true,
//				Pos:     p.VotePosCount(),
//				Neg:     p.VoteNegCount() - 1,
//			})
//		}
//
//		err := ushio.Data.PostNewNegVote(p.PID, nav.User.UID)
//		if err != nil {
//			return err
//		}
//
//		return ctx.JSON(&VoteResp{
//			Success: true,
//			Pos:     p.VotePosCount(),
//			Neg:     p.VoteNegCount() + 1,
//		})
//	}
//	return nil
//}

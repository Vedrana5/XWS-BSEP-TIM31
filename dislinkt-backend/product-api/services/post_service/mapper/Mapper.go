package mapper

import (
	"common/module/proto/post_service"
	"post/module/model"
)

func MapPostReply(post *model.Post) *post_service.Post {
	id := post.Id.Hex()

	likesNum, dislikesNum := FindNumberOfReactions(post)

	postPb := &post_service.Post{
		Id:             id,
		Username:       post.Username,
		PostText:       post.PostText,
		DatePosted:     post.DatePosted.String(),
		LikesNumber:    int32(likesNum),
		DislikesNumber: int32(dislikesNum),
		CommentsNumber: int32(len(post.Comments)),
	}
	//	postPb.ImagePaths = convertByteToBase64(post.ImagePaths)

	return postPb
}

func FindNumberOfReactions(post *model.Post) (int, int) {
	likesNum := 0
	dislikesNum := 0

	for _, reaction := range post.Reactions {
		if reaction.Reaction == model.LIKED {
			likesNum++
		} else if reaction.Reaction == model.DISLIKED {
			dislikesNum++
		}
	}
	return likesNum, dislikesNum
}

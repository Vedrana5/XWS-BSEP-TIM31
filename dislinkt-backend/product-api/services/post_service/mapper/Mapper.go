package mapper

import (
	"common/module/proto/post_service"
	b64 "encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"image"
	"image/jpeg"
	"log"
	"os"
	"post/module/model"
	"strings"
	"time"
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
	postPb.ImagePaths = convertByteToBase64(post.ImagePaths)

	return postPb
}

func convertByteToBase64(paths []byte) string {
	imageEnc := b64.StdEncoding.EncodeToString(paths)
	return imageEnc

}

func MapNewJobOffer(offerPb *post_service.JobOffer) *model.JobOffer {

	offer := &model.JobOffer{
		Id:             primitive.NewObjectID(),
		Publisher:      offerPb.Publisher,
		Position:       offerPb.Position,
		JobDescription: offerPb.JobDescription,
		Requirements:   offerPb.Requirements,
		DatePosted:     mapToDate(offerPb.DatePosted),
		Duration:       mapToDate(offerPb.Duration),
	}

	return offer
}

func mapToDate(posted string) time.Time {
	layout := "2006-01-02T15:04:05.000Z"
	dateOfBirth, _ := time.Parse(layout, posted)
	return dateOfBirth

}

func MapNewPost(postPb *post_service.Post) *model.Post {
	post := &model.Post{
		Id:         primitive.NewObjectID(),
		Username:   postPb.Username,
		PostText:   postPb.PostText,
		DatePosted: time.Now(),
		IsDeleted:  false,
	}
	base64toJpg(postPb.ImagePaths)
	post.ImagePaths = convertBase64ToByte(postPb.ImagePaths)
	return post
}

func MapUserCommentsForPost(userName string, commentText string) *post_service.Comment {

	commentPb := &post_service.Comment{
		Username:    userName,
		CommentText: commentText,
	}

	return commentPb
}

func MapJobOfferReply(offer *model.JobOffer) *post_service.JobOffer {
	id := offer.Id.Hex()

	offerPb := &post_service.JobOffer{
		Id:             id,
		Publisher:      offer.Publisher,
		Position:       offer.Position,
		JobDescription: offer.JobDescription,
		Requirements:   offer.Requirements,
		DatePosted:     offer.DatePosted.String(),
		Duration:       offer.Duration.String(),
	}

	return offerPb
}

func MapNewComment(commentPb *post_service.Comment) *model.Comment {
	comment := &model.Comment{
		Id:          primitive.NewObjectID(),
		Username:    commentPb.Username,
		CommentText: commentPb.CommentText,
	}
	return comment
}

func convertBase64ToByte(paths string) []byte {
	fmt.Println("convertBase64ToByte")
	imageDec := paths[strings.IndexByte(paths, ',')+1:]
	dec, err := b64.StdEncoding.DecodeString(imageDec)
	if err != nil {
		panic(err)
	}
	return dec
}

func base64toJpg(paths string) {
	data := paths[strings.IndexByte(paths, ',')+1:]
	reader := b64.NewDecoder(b64.StdEncoding, strings.NewReader(data))
	m, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()
	fmt.Println("base64toJpg", bounds, formatString)

	//Encode from image format to writer
	pngFilename := "test"
	f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 75})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Jpg file", pngFilename, "created")
	err = f.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

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

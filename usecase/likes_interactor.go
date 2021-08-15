package usecase

import "fmt"

var ErrNotFound = fmt.Errorf("not found")

type LikesInteractor interface {
	LikePost(postID, userID int) error
}

type LikesInteractorImpl struct {
	usersRepository UsersRepository
	postsRepository PostsRepository
	pushRepository  PushRepository
}

func NewLikesInteractor(ur UsersRepository, por PostsRepository, pur PushRepository) *LikesInteractorImpl {
	return &LikesInteractorImpl{
		usersRepository: ur,
		postsRepository: por,
		pushRepository:  pur,
	}
}

func (li *LikesInteractorImpl) LikePost(postID, userID int) error {
	// いいね
	if err := li.postsRepository.LikePost(postID, userID); err != nil {
		return err
	}

	// 通知を送信など
	user, err := li.usersRepository.Get(userID)
	if err != nil {
		return err
	}
	// 略
	li.pushRepository.MessageLike(postID, user.ID)
	// 略

	return nil
}

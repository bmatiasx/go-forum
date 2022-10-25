package repository

import (
	"errors"
	"fmt"

	"bitbucket.org/mindera/go-rest-blog/model"
)

type CommentRepository struct {
	repository []model.Comment
}

func NewCommentRepository() *CommentRepository {
	repo := CustomCommentRepository(make([]model.Comment, 0))
	return &repo
}

func CustomCommentRepository(mockStorage []model.Comment) CommentRepository {
	return CommentRepository{repository: mockStorage}
}

type CommentAlreadyExistsError struct {
	id uint64
}

func (e CommentAlreadyExistsError) Error() string {
	return fmt.Sprintf("Error: Comment with id: %v already exists in the repository!", e.id)
}

type CommentNotFoundError struct {
	id uint64
}

func (e CommentNotFoundError) Error() string {
	return fmt.Sprintf("Error: Comment with id: %v was not found in the repository!", e.id)
}

func (c *CommentRepository) Insert(comment model.Comment) error {
	// TODO: Insert should insert a comment passed as an argument to the persistent in memory repository.
	//  The method should return an apierror as an instance of `CommentAlreadyExistsError` struct
	//  when a comment with given id already exists in the repository.
	err := CommentAlreadyExistsError{}

	if len(c.repository) == 0 {
		c.repository = append(c.repository, comment)
		return nil
	}

	for _, cm := range c.repository {
		if comment.Id == cm.Id {
			err.id = cm.Id
			return errors.New(err.Error())
		}
		c.repository = append(c.repository, comment)
	}

	return nil
}

func (c *CommentRepository) GetById(id uint64) (*model.Comment, error) {
	// TODO: GetById should return a comment from a repository that has a given id.
	//  If there's no comment with given id, this function should return a (nil, CommentNotFoundError) pair
	//  with CommentNotFound instance having id member variable set with id passed to this method.
	var found *model.Comment

	for _, c := range c.repository {
		if c.Id == id {
			found = &c
			break
		}
	}

	if found == nil {
		err := CommentNotFoundError{
			id: id,
		}
		return nil, errors.New(err.Error())
	}

	return found, nil
}

func (c *CommentRepository) GetAllByPostId(id uint64) []model.Comment {
	// TODO: GetAllByPostId should return a slice of all comments that have PostId member variable
	//  equal to given id.
	//  The method should return an empty slice when there are no comments with given id in the repository.
	comments := make([]model.Comment, 0)

	for _, c := range c.repository {
		if c.PostId == id {
			comments = append(comments, c)
		}
	}
	return comments
}

type PostRepository struct {
	repository []model.Post
}

func CustomPostRepository(mockStorage []model.Post) PostRepository {
	return PostRepository{repository: mockStorage}
}

func NewPostRepository() *PostRepository {
	repo := CustomPostRepository(make([]model.Post, 0))
	return &repo
}

type PostAlreadyExistsError struct {
	id uint64
}

func (e PostAlreadyExistsError) Error() string {
	return fmt.Sprintf("Error: Post with id: %v already exists in the repository!", e.id)
}

type PostNotFoundError struct {
	id uint64
}

func (e PostNotFoundError) Error() string {
	return fmt.Sprintf("Error: Post with id: %v was not found in the repository!", e.id)
}

func (c *PostRepository) Insert(post model.Post) error {
	// TODO:  Insert should insert a post passed as an argument to the persistent in memory repository.
	//  The method should return an apierror as an instance of `PostAlreadyExistsError` struct
	//  when a post with given id already exists in the repository.
	err := PostAlreadyExistsError{}

	if len(c.repository) == 0 {
		c.repository = append(c.repository, post)
		return nil
	}

	for _, p := range c.repository {
		if post.Id == p.Id {
			err.id = post.Id
			return errors.New(err.Error())
		}
		c.repository = append(c.repository, post)
	}

	return nil
}

func (c *PostRepository) GetById(id uint64) (*model.Post, error) {
	// TODO: GetById should return a post from a repository that has a given id.
	//  If there's no post with given id, this function should return a (nil, PostNotFoundError) pair
	//  with PostNotFoundError instance having id member variable set with id passed to this method.
	var found *model.Post
	for _, c := range c.repository {
		if c.Id == id {
			found = &c
			break
		}
	}

	if found == nil {
		err := PostNotFoundError{
			id: id,
		}
		return nil, errors.New(err.Error())
	}

	return found, nil
}

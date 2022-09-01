package db

import (
	"context"
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) Post {
	user := createRandomUser(t)
	arg := CreatePostParams{
		Title: util.RandomStr(12),
		Description: util.RandomStr(20),
		Content: util.RandomStr(80),
		UserName: user.Username,
	}

	post, err := testQueries.CreatePost(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, user.Username, post.UserName)
	require.Equal(t, post.Title, arg.Title)
	require.Equal(t, post.Description, arg.Description)
	require.Equal(t, post.Content, arg.Content)
	require.NotZero(t, post.ID)

	return post;
}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestGetPost(t *testing.T) {
	post := createRandomPost(t)

	post2, err := testQueries.GetPost(context.Background(), post.ID)

	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post.Content, post2.Content)
	require.Equal(t, post.Title, post2.Title)
	require.Equal(t, post.Description, post2.Description)

	require.NotZero(t, post2.ID)
}

func TestListPosts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomPost(t)
	}
	posts, err := testQueries.ListPosts(context.Background(), ListPostsParams{
		Limit: 5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.NotEmpty(t, posts)
	require.Len(t, posts, 5)

	for _, post := range posts{
		require.NotEmpty(t, post)
	}
}

func TestUpdatePosts(t *testing.T) {
	post := createRandomPost(t)
	
	post1, err := testQueries.UpdatePost(context.Background(), UpdatePostParams{
		ID: post.ID,
		Content: util.RandomStr(100),
	})

	require.NoError(t, err)
	require.NotEmpty(t, post1)
	require.NotEqual(t, post.Content, post1.Content)
	require.Equal(t, post.Description, post1.Description)
	require.Equal(t, post.Title, post1.Title)
	require.Equal(t, post.UserName, post1.UserName)
}

func TestDeletePost(t *testing.T) {
	post := createRandomPost(t)

	err := testQueries.DeletePost(context.Background(), post.ID)

	require.NoError(t, err)
}

// func TestGetAllUsersPost(t *testing.T) {
	// var post Post
	// for i := 0; i < 8; i++ {
		// post = createRandomPost(t)
	// }

	// posts, err := testQueries.ListPostsForUser(context.Background(), ListPostsForUserParams{
		// Column1: post.UserName,
		// Limit: 5,
		// Offset: 5,
	// }) 

	// require.NoError(t, err)
	// require.NotEmpty(t, posts)

	// for _, list := range posts {
		// require.Equal(t, list.UserName, post.UserName)
		// require.Len(t, list, 5)
	// }
// }
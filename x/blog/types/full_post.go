package types

import "fmt"

func (storedPost StoredPost) GetPostTitle() (title string, err error) {
	if len(storedPost.Title) <= 0 {
		return title, ErrMissingPostTitle.Wrap(fmt.Sprintf("index = %s", storedPost.Index))
	}
	return title, nil
}

func (storedPost StoredPost) GetPostBody() (body string, err error) {
	if len(storedPost.Body) <= 0 {
		return body, ErrMissingPostBody.Wrap(fmt.Sprintf("index = %s", storedPost.Index))
	}
	return body, nil
}

func (storedPost StoredPost) Validate() (err error) {
	_, err = storedPost.GetPostTitle()
	if err != nil {
		return err
	}

	_, err = storedPost.GetPostBody()
	if err != nil {
		return err
	}
	
	return err
}
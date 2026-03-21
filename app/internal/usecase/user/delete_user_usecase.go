package usecase

import "fmt"

func (u *userUsecase) DeleteUser(loggedInUserId int, userIds []int) error {
	for _, userId := range userIds {
		_, err := u.userRepository.GetUserDetail(loggedInUserId, userId)
		if err != nil {
			return fmt.Errorf("id not found")
		}
	}

	return u.userRepository.DeleteUser(loggedInUserId, userIds)
}

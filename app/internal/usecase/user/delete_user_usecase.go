package usecase

import "fmt"

func (u *userUsecase) DeleteUser(loggedInUserId int, userIds []int) error {
	for _, userId := range userIds {
		_, err := u.userRepository.FindByID(userId)
		if err != nil {
			return fmt.Errorf("id not found")
		}
	}

	return u.userRepository.Delete(userIds)
}

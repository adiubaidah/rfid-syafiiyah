package db

import (
	"context"
	"testing"

	"github.com/adiubaidah/rfid-syafiiyah/internal/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func clearUserTable(t *testing.T) {
	_, err := testQueries.db.Exec(context.Background(), `DELETE FROM "user"`)
	require.NoError(t, err)
}

func createRandomUser(t *testing.T, role UserRole) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username: util.RandomString(8),
		Role:     role,
		Password: hashedPassword,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username.String)
	require.Equal(t, arg.Role, user.Role.UserRole)

	return user
}

func TestCreateUser(t *testing.T) {
	roles := []UserRole{UserRoleSuperadmin, UserRoleAdmin, UserRoleEmployee, UserRoleParent}
	n := len(roles)
	createRandomUser(t, roles[util.RandomInt(0, int64(n-1))])
}

func TestListUsersRelation(t *testing.T) {
	t.Run("Should return users matching List", func(t *testing.T) {
		clearUserTable(t)
		user1 := createRandomUser(t, UserRoleSuperadmin)
		createRandomUser(t, UserRoleAdmin)
		createRandomUser(t, UserRoleEmployee)
		createRandomUser(t, UserRoleParent)

		arg := ListUsersAscUsernameParams{
			Q:            pgtype.Text{String: user1.Username.String[:3], Valid: true},
			Role:         NullUserRole{Valid: false},
			LimitNumber:  10,
			OffsetNumber: 0,
			HasRelation:  -1,
		}

		users, err := testQueries.ListUsersAscUsername(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, users)
	})

	t.Run("Should return empty array when HasRelation is true but no relations exist", func(t *testing.T) {
		clearUserTable(t)
		createRandomUser(t, UserRoleSuperadmin)
		createRandomUser(t, UserRoleAdmin)
		createRandomUser(t, UserRoleEmployee)
		createRandomUser(t, UserRoleParent)

		arg := ListUsersAscUsernameParams{
			Q:            pgtype.Text{String: "", Valid: false},
			Role:         NullUserRole{Valid: false},
			LimitNumber:  10,
			OffsetNumber: 0,
			HasRelation:  1,
		}

		users, err := testQueries.ListUsersAscUsername(context.Background(), arg)
		require.NoError(t, err)
		require.Empty(t, users)
	})

	t.Run("Should return users with relations", func(t *testing.T) {
		clearUserTable(t)
		_, user := createRandomParentWithUser(t)
		createRandomUser(t, UserRoleSuperadmin)

		arg := ListUsersAscUsernameParams{
			Q:            pgtype.Text{String: "", Valid: false},
			Role:         NullUserRole{Valid: false},
			LimitNumber:  10,
			OffsetNumber: 0,
			HasRelation:  1,
		}
		users, err := testQueries.ListUsersAscUsername(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, users)
		require.Equal(t, user.Username.String, users[0].Username.String)
	})
}

func TestListUserPagination(t *testing.T) {
	clearUserTable(t)
	for i := 0; i < 10; i++ {
		createRandomUser(t, UserRoleSuperadmin)
	}

	testCases := []struct {
		name     string
		arg      ListUsersAscUsernameParams
		expected int
	}{
		{
			name: "Limit 5",
			arg: ListUsersAscUsernameParams{
				LimitNumber:  5,
				OffsetNumber: 0,
			},
			expected: 5,
		},
		{
			name: "Limit 5 Offset 5",
			arg: ListUsersAscUsernameParams{
				LimitNumber:  5,
				OffsetNumber: 5,
			},
			expected: 5,
		},
		{
			name: "Limit 5 Offset 10",
			arg: ListUsersAscUsernameParams{
				LimitNumber:  5,
				OffsetNumber: 10,
			},
			expected: 0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			employees, err := testQueries.ListUsersAscUsername(context.Background(), tt.arg)
			require.NoError(t, err)
			require.Len(t, employees, tt.expected)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	clearUserTable(t)
	user1 := createRandomUser(t, UserRoleSuperadmin)

	arg := UpdateUserParams{
		Username: pgtype.Text{String: util.RandomString(8), Valid: true},
		Role:     NullUserRole{Valid: true, UserRole: UserRoleAdmin},
		ID:       user1.ID,
	}

	user, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username.String, user.Username.String)
	require.Equal(t, arg.Role.UserRole, user.Role.UserRole)
}

func TestDeleteUser(t *testing.T) {
	clearUserTable(t)
	user1 := createRandomUser(t, UserRoleSuperadmin)

	userDeleted, err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, userDeleted)

	require.Equal(t, user1.ID, userDeleted.ID)
	require.Equal(t, user1.Username.String, userDeleted.Username.String)
	require.Equal(t, user1.Role.UserRole, userDeleted.Role.UserRole)
}

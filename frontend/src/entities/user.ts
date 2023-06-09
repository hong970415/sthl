export interface IUser {
  id: string
  email: string
  password: string
  emailVerified: boolean
  isArchived: boolean
  createdAt: string
  updatedAt: string
}

export type ISensitiveUser = Omit<IUser, 'password'>
//----
// enum IUserFields {
//   ID = 'id',
//   Email = 'email',
//   Password = 'password',
//   EmailVerified = 'emailVerified',
//   IsArchived = 'isArchived',
//   CreatedAt = 'createdAt',
//   UpdatedAt = 'updatedAt',
// }

// interface IRUser {
//   [IUserFields.ID]: string
//   [IUserFields.Email]: string
//   [IUserFields.Password]: string
//   [IUserFields.EmailVerified]: boolean
//   [IUserFields.IsArchived]: boolean
//   [IUserFields.CreatedAt]: string
//   [IUserFields.UpdatedAt]: string
// }
// type IRSensitiveUser = Omit<IRUser, IUserFields.Password>

// type ISignupDto = Pick<IRUser, IUserFields.Email | IUserFields.Password>
// type ILoginDto = Pick<IRUser, IUserFields.Email | IUserFields.Password>
// type IUserUpdatePwDto = {
//   currentPassword: string
//   newPassword: string
// }

// const a: IRUser = {
//   [IUserFields.ID]: '',
//   [IUserFields.Email]: '',
//   [IUserFields.Password]: '',
//   [IUserFields.EmailVerified]: false,
//   [IUserFields.IsArchived]: false,
//   [IUserFields.CreatedAt]: '',
//   [IUserFields.UpdatedAt]: '',
// }

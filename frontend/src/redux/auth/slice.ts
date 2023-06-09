import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { IPassport } from '@/services/auth'
import { thunkGetMe, thunkLogin } from './thunk'
import { ISensitiveUser } from '@/entities/user'

interface IAuthState {
  isLoggedIn: boolean
  isLoggingIn: boolean
  passport: IPassport
  user: ISensitiveUser | null
}
const initialState: IAuthState = {
  isLoggedIn: false,
  isLoggingIn: false,
  passport: {
    accessToken: '',
    refreshToken: '',
  },
  user: null,
}

export const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    setPassport: (state, action: PayloadAction<IPassport>) => {
      state.passport = action.payload
    },
    logout: (state) => {
      state.isLoggedIn = false
      state.passport.accessToken = ''
      state.passport.refreshToken = ''
      state.user = null
    },
  },
  extraReducers: (builder) => {
    // thunkLogin
    builder.addCase(thunkLogin.pending, (state, action) => {
      state.isLoggedIn = false
      state.isLoggingIn = true
    })
    builder.addCase(thunkLogin.fulfilled, (state, action) => {
      state.isLoggedIn = true
      state.isLoggingIn = false
      state.passport = action.payload.data
    })
    builder.addCase(thunkLogin.rejected, (state, action) => {
      state.isLoggedIn = false
      state.isLoggingIn = false
    })
    // thunkGetMe
    builder.addCase(thunkGetMe.pending, (state, action) => {
      state.isLoggedIn = false
      state.isLoggingIn = true
    })
    builder.addCase(thunkGetMe.fulfilled, (state, action) => {
      state.isLoggedIn = true
      state.isLoggingIn = false
      state.user = action.payload.data
    })
    builder.addCase(thunkGetMe.rejected, (state, action) => {
      state.isLoggedIn = false
      state.isLoggingIn = false
    })
  },
})
export const { setPassport, logout } = authSlice.actions

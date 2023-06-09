import { createAsyncThunk } from '@reduxjs/toolkit'
import { showNotification } from '@mantine/notifications'
import { API } from '@/services'
import { ILoginForm } from '@/services/auth'

// thunkLogin
export const thunkLogin = createAsyncThunk(
  'auth/login',
  async (payload: ILoginForm, thunkApi) => {
    const response = await API.postLogin(payload)
    if (!response.success) {
      return thunkApi.rejectWithValue(response.errorMsg)
    }
    return thunkApi.fulfillWithValue(response.data)
  }
)

// thunkGetMe
export const thunkGetMe = createAsyncThunk(
  'auth/getMe',
  async (_: void, thunkApi) => {
    const response = await API.getMe()
    if (!response.success) {
      return thunkApi.rejectWithValue(response.errorMsg)
    }
    return thunkApi.fulfillWithValue(response.data)
  }
)

import { combineReducers } from 'redux'
import { useDispatch } from 'react-redux'
import { Context, createWrapper } from 'next-redux-wrapper'
import { configureStore } from '@reduxjs/toolkit'
import { authSlice } from './auth/slice'
import { authListenerMiddleware } from './auth/middleware'

const rootReducer = combineReducers({
  auth: authSlice.reducer,
})
export const store = configureStore({
  reducer: rootReducer,
  devTools: true,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat([authListenerMiddleware.middleware]),
})
export const makeStore = (context: Context) => {
  return store
}
export const wrapper = createWrapper(makeStore)

export type ReduxState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export const useAppDispatch: () => AppDispatch = useDispatch

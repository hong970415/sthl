import { ReactNode, useReducer } from 'react'
import { IProduct } from '@/entities/product'
import { ISiteUi } from '@/entities/site'
import { getInitialState, UserSiteContext } from './context'
import { userSiteStateActionReducer } from './reducer'
import { createResetUiDataAction } from './actions'

// ----UserSiteProvider
interface IUserSiteProviderProps {
  userId: string
  products: IProduct[]
  userUiData: ISiteUi
  editMode: boolean
  children?: ReactNode
}
export default function UserSiteProvider(props: IUserSiteProviderProps) {
  const { userId, products, userUiData, editMode, children } = props
  const reducerInitialState = getInitialState(
    userId,
    products,
    userUiData,
    editMode
  )
  const [state, stateDispatch] = useReducer(
    userSiteStateActionReducer,
    reducerInitialState
  )

  const handleResetUiData = () => {
    stateDispatch(createResetUiDataAction(reducerInitialState.uiData))
  }

  const value = {
    userSiteState: state,
    userSiteStateDispatch: stateDispatch,
    resetUiData: handleResetUiData,
  }
  return (
    <UserSiteContext.Provider value={value}>
      {children}
    </UserSiteContext.Provider>
  )
}

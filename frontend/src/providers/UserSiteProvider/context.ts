import { createContext, Dispatch } from 'react'
import { IProduct } from '@/entities/product'
import { ISiteUi } from '@/entities/site'
import { IUserSiteStateAction } from './actions'

export interface IUserSiteState {
  userId: string
  cartItems: IProduct[]
  readonly productsData: IProduct[]
  uiData: ISiteUi
  readonly editMode: boolean
}
export const UserSiteContext = createContext<{
  userSiteState: IUserSiteState
  userSiteStateDispatch: Dispatch<IUserSiteStateAction>
  resetUiData: () => void
} | null>(null)

export function getInitialState(
  userId: string,
  products: IProduct[],
  uiData: ISiteUi,
  editMode: boolean
) {
  return {
    userId: userId,
    cartItems: [],
    productsData: products,
    uiData: uiData,
    editMode: editMode,
  }
}

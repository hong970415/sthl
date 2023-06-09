import { IProduct } from '@/entities/product'
import { ISiteUi } from '@/entities/site'
import { createAction } from '@reduxjs/toolkit'

enum UserSiteStateActionType {
  UpdateSitename = 'updateSitename',
  UpdateHomepageImgUrl = 'updateHomepageImgUrl',
  UpdateHomepageText = 'updateHomepageText',
  UpdateHomepageTextColor = 'updateHomepageTextColor',
  ResetUiData = 'resetUiData',
  AddCartItem = 'addCartItem',
  RemoveCartItemById = 'removeCartItemById',
  IncrementCartItemQuantityById = 'incrementCartItemQuantityById',
  DecrementCartItemQuantityById = 'decrementCartItemQuantityById',
}

type IUserSiteStateUiAction =
  | { type: UserSiteStateActionType.UpdateSitename; payload: string }
  | { type: UserSiteStateActionType.UpdateHomepageImgUrl; payload: string }
  | { type: UserSiteStateActionType.UpdateHomepageText; payload: string }
  | { type: UserSiteStateActionType.UpdateHomepageTextColor; payload: string }
  | { type: UserSiteStateActionType.ResetUiData; payload: ISiteUi }
type IUserSiteStateAction =
  | IUserSiteStateUiAction
  | { type: UserSiteStateActionType.AddCartItem; payload: IProduct }
  | { type: UserSiteStateActionType.RemoveCartItemById; payload: string }
  | {
      type: UserSiteStateActionType.IncrementCartItemQuantityById
      payload: string
    }
  | {
      type: UserSiteStateActionType.DecrementCartItemQuantityById
      payload: string
    }

const createUpdateSitenameAction = createAction<
  string,
  UserSiteStateActionType.UpdateSitename
>(UserSiteStateActionType.UpdateSitename)

const createUpdateHomepageImgUrlAction = createAction<
  string,
  UserSiteStateActionType.UpdateHomepageImgUrl
>(UserSiteStateActionType.UpdateHomepageImgUrl)

const createUpdateHomepageTextAction = createAction<
  string,
  UserSiteStateActionType.UpdateHomepageText
>(UserSiteStateActionType.UpdateHomepageText)

const createUpdateHomepageTextColorAction = createAction<
  string,
  UserSiteStateActionType.UpdateHomepageTextColor
>(UserSiteStateActionType.UpdateHomepageTextColor)

const createResetUiDataAction = createAction<
  ISiteUi,
  UserSiteStateActionType.ResetUiData
>(UserSiteStateActionType.ResetUiData)

const createAddCartItemAction = createAction<
  IProduct,
  UserSiteStateActionType.AddCartItem
>(UserSiteStateActionType.AddCartItem)

const createRemoveCartItemByIdAction = createAction<
  string,
  UserSiteStateActionType.RemoveCartItemById
>(UserSiteStateActionType.RemoveCartItemById)

const createIncrementCartItemQuantityByIdAction = createAction<
  string,
  UserSiteStateActionType.IncrementCartItemQuantityById
>(UserSiteStateActionType.IncrementCartItemQuantityById)

const createDecrementCartItemQuantityByIdAction = createAction<
  string,
  UserSiteStateActionType.DecrementCartItemQuantityById
>(UserSiteStateActionType.DecrementCartItemQuantityById)

export type { IUserSiteStateUiAction, IUserSiteStateAction }
export {
  UserSiteStateActionType,
  createUpdateSitenameAction,
  createUpdateHomepageImgUrlAction,
  createUpdateHomepageTextAction,
  createUpdateHomepageTextColorAction,
  createResetUiDataAction,
  createAddCartItemAction,
  createRemoveCartItemByIdAction,
  createIncrementCartItemQuantityByIdAction,
  createDecrementCartItemQuantityByIdAction,
}

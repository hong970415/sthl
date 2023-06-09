import { IUserSiteState } from './context'
import { IUserSiteStateAction, UserSiteStateActionType } from './actions'

export function userSiteStateActionReducer(
  state: IUserSiteState,
  action: IUserSiteStateAction
): IUserSiteState {
  switch (action.type) {
    case UserSiteStateActionType.AddCartItem: {
      const isExisted =
        state.cartItems.findIndex((el) => el.id === action.payload.id) > -1
      if (isExisted) {
        return {
          ...state,
        }
      }
      return {
        ...state,
        cartItems: [...state.cartItems, action.payload],
      }
    }
    case UserSiteStateActionType.RemoveCartItemById: {
      return {
        ...state,
        cartItems: state.cartItems.filter((el) => el.id != action.payload),
      }
    }
    case UserSiteStateActionType.IncrementCartItemQuantityById: {
      return {
        ...state,
        cartItems: state.cartItems.map((el) => {
          if (el.id === action.payload) {
            return { ...el, quantity: el.quantity + 1 }
          }
          return { ...el }
        }),
      }
    }
    case UserSiteStateActionType.DecrementCartItemQuantityById: {
      return {
        ...state,
        cartItems: state.cartItems.map((el) => {
          if (el.id === action.payload) {
            return { ...el, quantity: el.quantity - 1 }
          }
          return { ...el }
        }),
      }
    }
    // ui
    case UserSiteStateActionType.UpdateSitename: {
      return {
        ...state,
        uiData: {
          ...state.uiData,
          sitename: action.payload,
        },
      }
    }
    case UserSiteStateActionType.UpdateHomepageImgUrl: {
      return {
        ...state,
        uiData: {
          ...state.uiData,
          homepageImgUrl: action.payload,
        },
      }
    }
    case UserSiteStateActionType.UpdateHomepageText: {
      return {
        ...state,
        uiData: {
          ...state.uiData,
          homepageText: action.payload,
        },
      }
    }
    case UserSiteStateActionType.UpdateHomepageTextColor: {
      return {
        ...state,
        uiData: {
          ...state.uiData,
          homepageTextColor: action.payload,
        },
      }
    }
    case UserSiteStateActionType.ResetUiData: {
      return {
        ...state,
        uiData: action.payload,
      }
    }
    default: {
      throw new Error(`Unhandled action type: ${JSON.stringify(action)}`)
    }
  }
}

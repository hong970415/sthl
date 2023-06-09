export type { IPassport, ILoginForm, IRefreshTokenForm } from './auth'
export type { ISignupForm, IUserUpdatePwForm } from './user'
export type { ICreateProductForm, IUpdateProductForm } from './product'
export type {
  ICreateOrderItem,
  ICreateOrderForm,
  IUpdateOrderForm,
} from './order'
export type {
  IServerResponse,
  IPagingQuery,
  IPagingQueryResult,
  IServerResponseWithPaging,
  CustomAxiosResponse,
} from './httpClient'

import { StatusCode } from './httpClient'
import { postLogin, postRefreshToken } from './auth'
import {
  getPing,
  getMe,
  postSignup,
  putUpdatePassword,
  checkUserExistById,
} from './user'
import {
  postCreateProduct,
  getProducts,
  getProductById,
  putUpdateProductById,
  deleteProductById,
} from './product'
import {
  postCreateOrder,
  getOrders,
  getOrderById,
  putUpdateOrderById,
  deleteOrderById,
} from './order'
import { getUserSiteUiDataById, putUpsertUserSiteUiDataById } from './site'
import {
  postUploadAlbumImg,
  getAlbumImgsByUserId,
  putUpdateImgDataById,
} from './album'
const API = {
  // auth
  postLogin,
  postRefreshToken,
  // user
  getPing,
  getMe,
  postSignup,
  putUpdatePassword,
  checkUserExistById,
  // product,
  postCreateProduct,
  getProducts,
  getProductById,
  putUpdateProductById,
  deleteProductById,
  // order
  postCreateOrder,
  getOrders,
  getOrderById,
  putUpdateOrderById,
  deleteOrderById,
  // site
  getUserSiteUiDataById,
  putUpsertUserSiteUiDataById,
  // album
  postUploadAlbumImg,
  getAlbumImgsByUserId,
  putUpdateImgDataById,
}

export { API, StatusCode }

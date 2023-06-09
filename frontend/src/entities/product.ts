export const ProductFieldKeys = [
  'id',
  'name',
  'price',
  'quantity',
  'description',
  'status',
  'createdAt',
  'updatedAt',
]
export interface IProduct {
  id: string
  createdAt: string
  updatedAt: string
  userId: string
  name: string
  price: number
  quantity: number
  description: string
  status: string
  isArchived: boolean
  imgUrl: string
}

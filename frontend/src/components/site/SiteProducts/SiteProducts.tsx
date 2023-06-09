import 'swiper/css'
import 'swiper/css/pagination'
import 'swiper/css/navigation'
import { MouseEvent } from 'react'
import { Swiper, SwiperSlide } from 'swiper/react'
import { Navigation, Pagination } from 'swiper'
import { Box, Button, Group, Paper, Stack, Text } from '@mantine/core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCircleCheck } from '@fortawesome/free-regular-svg-icons'
import { IProduct } from '@/entities/product'
import { getPrice } from '@/utils/price'
import useUserSite from '@/hooks/userUserSite'
import useStyles from './SiteProducts.style'
import {
  createAddCartItemAction,
  createRemoveCartItemByIdAction,
} from '@/providers/UserSiteProvider/actions'
interface ISiteProductsProps {
  id: string
  products: IProduct[]
}
export default function SiteProducts(props: ISiteProductsProps) {
  const { id, products = [] } = props
  const { classes } = useStyles({})
  const { userSiteState, userSiteStateDispatch } = useUserSite()

  const handleOnClickAddToCart = (
    event: MouseEvent,
    isAddedToCart: boolean,
    value: IProduct
  ) => {
    if (isAddedToCart) {
      userSiteStateDispatch(createRemoveCartItemByIdAction(value.id))
      return
    }
    const payload = {
      ...value,
      quantity: 1,
    }
    userSiteStateDispatch(createAddCartItemAction(payload))
  }

  const data = products
  console.log('data', data)

  const slidesEl =
    data &&
    data.map((productItem) => {
      const isAddedToCart =
        userSiteState.cartItems.findIndex((el) => el.id === productItem.id) > -1
      return (
        <SwiperSlide key={productItem.id} style={{ height: 'auto' }}>
          <Paper className={classes.card} withBorder>
            <img src={productItem.imgUrl} className={classes.cardImg} />
            <Stack p={'lg'} justify="space-between" sx={{ flexGrow: 1 }}>
              <Stack>
                <Group position="apart">
                  <Text fw={500} fz="lg">
                    {productItem.name}
                  </Text>
                </Group>
                <Text fz="sm" c="dimmed">
                  {productItem.description}
                </Text>
              </Stack>

              <Group position="apart" noWrap>
                <div>
                  <Text fz="xl" span fw={500} className={classes.cardPrice}>
                    ${getPrice(productItem.price)}
                  </Text>
                </div>
                <Button
                  radius="md"
                  onClick={(event) =>
                    handleOnClickAddToCart(event, isAddedToCart, productItem)
                  }
                  color={isAddedToCart ? 'green' : 'blue'}
                  leftIcon={
                    isAddedToCart && <FontAwesomeIcon icon={faCircleCheck} />
                  }
                >
                  {isAddedToCart ? 'Added to cart' : 'Add to cart'}
                </Button>
              </Group>
            </Stack>
          </Paper>
        </SwiperSlide>
      )
    })
  return (
    <Box id={id} p="xl">
      <Swiper
        modules={[Navigation, Pagination]}
        spaceBetween={40}
        // onSlideChange={() => console.log('slide change')}
        // onSwiper={(swiper) => console.log(swiper)}
        navigation={{ enabled: true }}
        pagination={{ clickable: true }}
        breakpoints={{
          640: {
            slidesPerView: 2,
            spaceBetween: 20,
            navigation: { enabled: true },
          },
          768: {
            slidesPerView: 3,
            spaceBetween: 30,
            navigation: { enabled: true },
          },
          1200: {
            slidesPerView: 4,
            spaceBetween: 50,
            navigation: { enabled: true },
          },
        }}
        style={{ paddingBottom: '40px' }}
      >
        {slidesEl}
      </Swiper>
    </Box>
  )
}

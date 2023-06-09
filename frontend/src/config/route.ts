import {
  faArrowRightFromBracket,
  faBagShopping,
  faFileLines,
  faGear,
  faHouse,
  faImages,
  faShop,
} from '@fortawesome/free-solid-svg-icons'

export const PublicRoutes = ['/', '/404', '/500', '/cms/login', '/site/[id]']
export function getMainRoutes() {
  return [
    {
      group: 'dashboard',
      route: '/cms/dashboard',
      labelKey: 'sidebar.dashboard',
      icon: faHouse,
    },
    {
      group: 'album',
      route: '/cms/album',
      labelKey: 'sidebar.album',
      icon: faImages,
    },
    {
      group: 'sitedesign',
      route: '/cms/sitedesign',
      labelKey: 'sidebar.sitedesign',
      icon: faShop,
    },
    {
      group: 'product',
      route: '/cms/product/list',
      labelKey: 'sidebar.product',
      icon: faBagShopping,
    },
    {
      group: 'order',
      route: '/cms/order/list',
      labelKey: 'sidebar.order',
      icon: faFileLines,
    },
  ]
}
export function getSettingRoutes() {
  return [
    {
      group: 'setting',
      route: '/cms/setting',
      labelKey: 'sidebar.setting',
      icon: faGear,
      withLink: true,
    },
    {
      group: 'logout',
      route: '/cms/logout',
      labelKey: 'general.logout',
      icon: faArrowRightFromBracket,
      withLink: false,
    },
  ]
}

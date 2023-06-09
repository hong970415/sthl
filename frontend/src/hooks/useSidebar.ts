import { useState } from 'react'

export interface IUseSidebarState {
  status: 'opened' | 'closed' | 'minimizied'
}

/** useSidebar
 * @returns {boolean} isCollapsed - Sidebar open or not
 * @returns {number} width - Navbar closed or not
 * @returns {function} toggleSidebar - Toggle sidebar
 */
const SIDE_BAR_WIDTH = 200
export default function useSidebar() {
  const [state, setState] = useState({
    isCollapsed: true,
    width: 0,
    // isCollapsed: false,
    // width: SIDE_BAR_WIDTH,
  })

  const toggleSidebar = () =>
    setState((prev) => {
      if (prev.isCollapsed) {
        return {
          ...prev,
          isCollapsed: !prev.isCollapsed,
          width: SIDE_BAR_WIDTH,
        }
      }
      return { ...prev, isCollapsed: !prev.isCollapsed, width: 0 }
    })

  return {
    ...state,
    toggleSidebar: toggleSidebar,
  }
}

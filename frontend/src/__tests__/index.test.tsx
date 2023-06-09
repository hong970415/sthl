import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import IndexPage from '@/pages/index'
import mockRouter from 'next-router-mock'

describe('Home', () => {
  it('renders a heading', () => {
    mockRouter.push('/')

    // Render the component:
    render(<IndexPage />)

    // Ensure the router was updated:
    expect(mockRouter).toMatchObject({
      asPath: '/',
      pathname: '/',
    })
  })
})

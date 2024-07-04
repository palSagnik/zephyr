import React, { ReactNode } from 'react'
import Navbar from '../components/Navbar'

interface HomePageLayoutProps {
  children: ReactNode;
}

export default function HomePageLayout({ children }: HomePageLayoutProps) {
  return (
    <>
        <Navbar />
        {children}
    </>
  )
}

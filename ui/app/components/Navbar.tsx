import Link from 'next/link'
import React from 'react'

export default function Navbar() {
    return (
        <nav>
            <h2 className='text-center'>Zephyr</h2>
            <Link href="/auth/signup">Register</Link>
            <Link href="/auth/login">Login</Link>
        </nav>
    )
}

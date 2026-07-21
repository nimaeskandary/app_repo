import MuiContainer from '@mui/material/Container'
import React from 'react'

export type ContainerProps = {
    children: React.ReactNode
}

export const Container = ({children}: ContainerProps) => {
    return <MuiContainer>
        {children}
    </MuiContainer>
}

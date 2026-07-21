import MuiGrid from '@mui/material/Grid'
import React from 'react'

export type GridProps = {
  children: React.ReactNode
  columns?: number
}

export const Grid = ({ children, columns }: GridProps): React.JSX.Element => {
  return <MuiGrid container columns={columns}>
    {children}
  </MuiGrid>
}

export type GridItemProps = {
  children: React.ReactNode
  size: number
}

export const GridItem = ({children, size}: GridItemProps): React.JSX.Element => {
  return <MuiGrid size={size}>
    {children}
  </MuiGrid>
}

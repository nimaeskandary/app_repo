import MuiBox from '@mui/material/Box'
import React from 'react'

export type BoxProps = {
  children: React.ReactNode
}

export const Box = ({ children }: BoxProps): React.JSX.Element => {
  return <MuiBox>{children}</MuiBox>
}

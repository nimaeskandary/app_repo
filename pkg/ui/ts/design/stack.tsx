import MuiStack from '@mui/material/Stack'
import React from 'react'

export type StackProps = {
  children: React.ReactNode
}

export const Stack = ({ children }: StackProps): React.JSX.Element => {
  return <MuiStack>{children}</MuiStack>
}

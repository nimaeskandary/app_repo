import MuiTypography from '@mui/material/Typography'
import React from 'react'

export type TypographyProps = {
  children: React.ReactNode
}

const typographSx = {
  // prevent horizontal scroll bars if a word is too big to fit on a mobile screen
  overflowWrap: 'anywhere' 
}

export const H1 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography
    variant="h1"
    sx={typographSx}>
    {children}
    </MuiTypography>
}

export const H2 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography 
    variant="h2"
    sx={typographSx}>
    {children}
    </MuiTypography>
}

export const H3 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography 
    variant="h3"
    sx={typographSx}>
    {children}
    </MuiTypography>
}

export const H4 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography 
    variant="h4"
    sx={typographSx}>
    {children}
    </MuiTypography>
}

export const H5 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography 
    variant="h5"
    sx={typographSx}>
    {children}
    </MuiTypography>
}

export const H6 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography 
    variant="h6"
    sx={typographSx}>
    {children}
    </MuiTypography>
}

export const Subtitle1 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography 
    variant="subtitle1"
    sx={typographSx}>
    {children}
    </MuiTypography>
}

export const Subtitle2 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography
    variant="subtitle2"
    sx={typographSx}>
    {children}
  </MuiTypography>
}

export const Body1 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="body1"
    sx={typographSx}>
    {children}
    </MuiTypography>
}

export const Body2 = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="body2"
    sx={typographSx}>
    {children}
    </MuiTypography>
}

export const ButtonText = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="button"
    sx={typographSx}>
    {children}
  </MuiTypography>
}

export const Caption = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="caption"
    sx={typographSx}>
    {children}
  </MuiTypography>
}

export const Overline = ({ children }: TypographyProps): React.JSX.Element => {
  return <MuiTypography variant="overline"
    sx={typographSx}>
    {children}
  </MuiTypography>
}

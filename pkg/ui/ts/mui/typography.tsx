import MuiTypography, { type TypographyProps as MuiTypographyProps } from '@mui/material/Typography'
import type { SxProps, Theme } from '@mui/material/styles'

export type TypographyProps = Omit<MuiTypographyProps, 'variant'>

const typographSx: SxProps<Theme> = {
  // prevent horizontal scroll bars if a word is too big to fit on a mobile screen
  overflowWrap: 'anywhere' 
}

const createTypography = (variant: MuiTypographyProps['variant']) => ({ sx, ...props }: TypographyProps) => {
  const mergedSx = sx ? [typographSx, sx] as SxProps<Theme> : typographSx

  return <MuiTypography {...props} variant={variant} sx={mergedSx} />
}

export const H1 = createTypography('h1')
export const H2 = createTypography('h2')
export const H3 = createTypography('h3')
export const H4 = createTypography('h4')
export const H5 = createTypography('h5')
export const H6 = createTypography('h6')
export const Subtitle1 = createTypography('subtitle1')
export const Subtitle2 = createTypography('subtitle2')
export const Body1 = createTypography('body1')
export const Body2 = createTypography('body2')
export const ButtonText = createTypography('button')
export const Caption = createTypography('caption')
export const Overline = createTypography('overline')

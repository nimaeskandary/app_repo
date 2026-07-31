import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import GordlePage from '@app_repo/pkg_gordle_frontend/views/gordle/GordlePage'

beforeEach(() => {
  vi.spyOn(Math, 'random').mockReturnValue(0)
})

afterEach(() => {
  vi.restoreAllMocks()
})

describe('GordlePage', () => {
  it('should complete the game after the user commits the rubric word', async () => {
    const user = userEvent.setup()
    render(<GordlePage />)
    const commitButton = screen.getByTestId('commit-word')

    expect(commitButton).toBeDisabled()

    for (const letter of 'CRANE') {
      await user.click(screen.getByTestId(`select-letter-${letter}`))
    }

    expect(commitButton).toBeEnabled()
    expect(screen.getByTestId('staged-word')).toHaveTextContent('CRANE')

    await user.click(commitButton)

    expect(screen.getByTestId('game-won')).toHaveTextContent('YOU WON!')
    expect(screen.queryByTestId('commit-word')).not.toBeInTheDocument()
  })
})

import { describe, expect, it } from 'vitest'
import { CellState, type Cell } from '@app_repo/pkg_gordle/domain/Cell'
import { createBoard, scoreGuess } from '@app_repo/pkg_gordle/domain/Game'

describe('createBoard', () => {
  it('should create a six-by-five blank board', () => {
    const blankCell: Cell = {
      letter: '',
      state: CellState.Unguessed,
    }

    expect(createBoard()).toEqual(
      Array.from({ length: 6 }, () =>
        Array.from({ length: 5 }, () => blankCell),
      ),
    )
  })
})

describe('scoreGuess', () => {
  it('should score correct, row-correct, and wrong letters', () => {
    expect(scoreGuess('CATER', 'CRANE')).toEqual([
      { letter: 'C', state: CellState.Correct },
      { letter: 'A', state: CellState.RowCorrect },
      { letter: 'T', state: CellState.Wrong },
      { letter: 'E', state: CellState.RowCorrect },
      { letter: 'R', state: CellState.RowCorrect },
    ])
  })

  it('should not score duplicate letters more times than they occur in the answer', () => {
    expect(scoreGuess('CREEP', 'CRANE')).toEqual([
      { letter: 'C', state: CellState.Correct },
      { letter: 'R', state: CellState.Correct },
      { letter: 'E', state: CellState.RowCorrect },
      { letter: 'E', state: CellState.Wrong },
      { letter: 'P', state: CellState.Wrong },
    ])
  })
})

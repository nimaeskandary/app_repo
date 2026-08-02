import { act, renderHook } from '@testing-library/react'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { CellState } from '@app_repo/pkg_gordle/domain/Cell'
import { useGordleViewModel } from '@app_repo/pkg_gordle/pages/gordle/useGordleViewModel'

beforeEach(() => {
  vi.spyOn(Math, 'random').mockReturnValue(0)
})

afterEach(() => {
  vi.restoreAllMocks()
})

describe('useGordleViewModel', () => {
  describe('initialization', () => {
    it('should create an incomplete game', () => {
      const { result } = renderHook(() => useGordleViewModel())

      expect(result.current.isGameComplete).toBe(false)
    })
  })

  describe('commitWord', () => {
    it('should commit accepted words to consecutive rows', () => {
      const { result } = renderHook(() => useGordleViewModel())

      act(() => {
        expect(result.current.commitWord('SLATE')).toBe(true)
      })
      act(() => {
        expect(result.current.commitWord('BRICK')).toBe(true)
      })

      expect(result.current.cells[0].map((cell) => cell.letter)).toEqual([
        'S',
        'L',
        'A',
        'T',
        'E',
      ])
      expect(result.current.cells[1].map((cell) => cell.letter)).toEqual([
        'B',
        'R',
        'I',
        'C',
        'K',
      ])
      expect(result.current.cells[2].every((cell) => !cell.letter)).toBe(true)
    })

    it('should reject a word when the board is full', () => {
      const { result } = renderHook(() => useGordleViewModel())
      const committedWords = [
        'SLATE',
        'BRICK',
        'CLOUD',
        'GHOST',
        'PLANT',
        'SHORE',
      ]

      committedWords.forEach((word) => {
        act(() => {
          expect(result.current.commitWord(word)).toBe(true)
        })
      })
      act(() => {
        expect(result.current.commitWord('MIGHT')).toBe(false)
      })

      expect(
        result.current.cells.map((row) =>
          row.map((cell) => cell.letter).join(''),
        ),
      ).toEqual(committedWords)
    })

    it('should complete the game and reject further words after committing the rubric', () => {
      const { result } = renderHook(() => useGordleViewModel())

      act(() => {
        expect(result.current.commitWord('CRANE')).toBe(true)
      })

      expect(result.current.isGameComplete).toBe(true)
      expect(result.current.cells[0].every(
        (cell) => cell.state === CellState.Correct,
      )).toBe(true)

      act(() => {
        expect(result.current.commitWord('SLATE')).toBe(false)
      })
      expect(result.current.cells[1].every((cell) => !cell.letter)).toBe(true)
    })
  })
})

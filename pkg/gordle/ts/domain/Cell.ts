export enum CellState {
  Unguessed = 'Unguessed',
  Wrong = 'Wrong',
  RowCorrect = 'RowCorrect',
  Correct = 'Correct',
}

export type Cell = {
  letter: string
  state: CellState
}

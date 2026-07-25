export enum GordleCellState {
  Unguessed = 'Unguessed',
  Wrong = 'Wrong',
  RowCorrect = 'RowCorrect',
  Correct = 'Correct',
}

export type GordleCell = {
  letter: string
  state: GordleCellState
}

import { useEffect, useRef } from 'react'

export type GreetViewProps = {
  name: string
  titleName: string
  time: string
  toastMessage: string
  isToastVisible: boolean
  onNameChange: (name: string) => void
  onGreet: () => void
}

// Show the actual Wails version this project was generated against.
const wailsVersion = 'v3.0.0-alpha2.117'

function GreetView({
  name,
  titleName,
  time,
  toastMessage,
  isToastVisible,
  onNameChange,
  onGreet,
}: GreetViewProps) {
  const titleNameRef = useRef<HTMLSpanElement | null>(null)

  useEffect(() => {
    const titleNameElement = titleNameRef.current
    if (!titleNameElement) {
      return
    }
    const current = titleNameElement.querySelector('.title-name-text:not(.is-outgoing)')
    if (!current || current.textContent === titleName) {
      return
    }
    const incoming = document.createElement('span')
    incoming.className = 'title-name-text is-entering'
    incoming.textContent = titleName
    current.classList.add('is-outgoing')
    titleNameElement.appendChild(incoming)
    // Force a reflow so the transitions run from the starting state.
    void incoming.offsetWidth
    incoming.classList.remove('is-entering')
    current.classList.add('is-leaving')
    current.addEventListener('transitionend', () => current.remove(), { once: true })
  }, [titleName])

  return (
    <>
      <main className="container">
        <header className="brand">
          <a className="brand-mark" data-wml-openURL="https://v3.wails.io" aria-label="Wails website">
            <img src="/wails.png" className="brand-logo" alt="Wails logo" />
          </a>
          <a className="brand-badge" data-wml-openURL="https://reactjs.org" aria-label="React">
            <img src="/react.svg" alt="React logo" />
          </a>
        </header>

        <h1 className="title">
          <span className="title-accent">Wails +</span>{' '}
          <span className="title-name" ref={titleNameRef}>
            <span className="title-name-text">React</span>
          </span>
        </h1>
        <p className="subtitle">Build beautiful cross-platform apps with Go and React.</p>

        <div className="greet">
          <div className="input-box">
            <svg className="input-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" aria-hidden="true"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" /><circle cx="12" cy="7" r="4" /></svg>
            <input aria-label="input" className="input" value={name} onChange={(e) => onNameChange(e.target.value)} type="text" placeholder="Your name" autoComplete="off" />
            <button aria-label="greet-btn" className="btn" onClick={onGreet}>Greet
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round" aria-hidden="true"><line x1="5" y1="12" x2="19" y2="12" /><polyline points="12 5 19 12 12 19" /></svg>
            </button>
          </div>
        </div>
      </main>

      <hr className="footer-divider" />
      <footer className="footer">
        <span className="footer-version"><span>{wailsVersion}</span></span>
        <span className="footer-time">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" aria-hidden="true"><circle cx="12" cy="12" r="10" /><polyline points="12 6 12 12 16 14" /></svg>
          <span>{time}</span>
        </span>
        <a className="footer-docs" data-wml-openURL="https://v3.wails.io" aria-label="Wails documentation">Docs
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" aria-hidden="true"><line x1="7" y1="17" x2="17" y2="7" /><polyline points="7 7 17 7 17 17" /></svg>
        </a>
      </footer>

      <div className={`toast${isToastVisible ? ' is-visible' : ''}`} role="status" aria-live="polite">
        <span className="toast-label">From Go</span>
        <span aria-label="result" className="toast-msg">{toastMessage}</span>
      </div>
    </>
  )
}

export default GreetView

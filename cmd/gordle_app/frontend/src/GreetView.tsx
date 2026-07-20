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
      <main className="mx-auto flex w-full max-w-135 flex-[1_0_auto] flex-col items-center justify-center text-center sm:mx-0 sm:items-start sm:text-left">
        <header className="mb-(--s-4) flex items-center gap-(--s-2)">
          <a className="brand-mark inline-flex cursor-pointer items-center justify-center transition-[transform,filter] duration-200 hover:-translate-y-0.5 hover:drop-shadow-[0_0_1em_rgba(255,77,77,0.45)]" data-wml-openURL="https://v3.wails.io" aria-label="Wails website">
            <img src="/wails.png" className="block h-(--s-5) w-auto object-contain" alt="Wails logo" />
          </a>
          <a className="brand-badge inline-flex cursor-pointer items-center justify-center transition-[transform,filter] duration-200 hover:-translate-y-0.5 hover:drop-shadow-[0_0_1em_rgba(255,77,77,0.45)]" data-wml-openURL="https://reactjs.org" aria-label="React">
            <img src="/react.svg" className="block h-[calc(var(--s-5)/1.272)] w-auto object-contain" alt="React logo" />
          </a>
        </header>

        <h1 className="m-0 mb-(--s-2) max-w-full text-[clamp(2.5rem,7.5vw,3.7rem)] leading-[1.05] font-extrabold tracking-[-0.02em] text-(--text) wrap-break-word">
          <span className="block bg-clip-text text-transparent [background-image:var(--grad)] [-webkit-text-fill-color:transparent] sm:inline">Wails +</span>{' '}
          <span className="title-name relative inline-block" ref={titleNameRef}>
            <span className="title-name-text">React</span>
          </span>
        </h1>
        <p className="m-0 mx-auto mb-(--s-5) max-w-lg text-[clamp(1.05rem,2.9vw,1.2rem)] leading-[1.55] text-(--muted) sm:mx-0">Build beautiful cross-platform apps with Go and React.</p>

        <div className="w-[86%] sm:w-[62%]">
          <div className="input-box flex w-full items-center gap-1.5 rounded-(--radius) border border-(--glass-border) bg-(--glass) py-0.75 pr-1 pl-2.5 backdrop-blur-md transition-colors duration-200 focus-within:border-[rgba(255,77,77,0.55)] sm:gap-1.75 sm:py-1 sm:pr-1.25 sm:pl-2.75">
            <svg className="size-5.25 flex-none text-(--muted) sm:size-5.75" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" aria-hidden="true"><path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" /><circle cx="12" cy="7" r="4" /></svg>
            <input aria-label="input" className="input min-w-0 flex-1 border-0 bg-transparent px-0.5 py-1 text-base text-(--text) outline-none placeholder:text-[rgba(154,166,192,0.7)] sm:py-1.25" value={name} onChange={(e) => onNameChange(e.target.value)} type="text" placeholder="Your name" autoComplete="off" />
            <button aria-label="greet-btn" className="btn inline-flex flex-none cursor-pointer items-center gap-1.25 rounded-lg border-0 px-3.25 py-1.25 text-[0.875rem] font-semibold text-white shadow-[0_0.375rem_1.125rem_rgba(255,45,114,0.35)] [background-image:var(--grad)] [transition:transform_0.1s_ease,box-shadow_0.2s_ease,opacity_0.2s_ease] hover:shadow-[0_0.5rem_1.5rem_rgba(255,45,114,0.5)] active:scale-[0.97] sm:px-3.75 sm:py-1.5" onClick={onGreet}>Greet
              <svg className="size-4 sm:size-4.25" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.2" strokeLinecap="round" strokeLinejoin="round" aria-hidden="true"><line x1="5" y1="12" x2="19" y2="12" /><polyline points="12 5 19 12 12 19" /></svg>
            </button>
          </div>
        </div>
      </main>

      <hr className="m-0 h-0 w-full border-0 border-t border-t-white/10" />
      <footer className="grid w-full grid-cols-[1fr_auto_1fr] items-center gap-1 pt-2.75 pr-[max(0.5rem,env(safe-area-inset-right))] pb-[max(0.6875rem,env(safe-area-inset-bottom))] pl-0 text-[0.6875rem] text-(--muted) sm:gap-0 sm:text-[0.78125rem]">
        <span className="inline-flex items-center justify-self-start gap-1.75 whitespace-nowrap"><span>{wailsVersion}</span></span>
        <span className="inline-flex items-center justify-self-center gap-1.25 sm:gap-1.75">
          <svg className="size-3.25 opacity-[0.85] sm:size-4.75" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" aria-hidden="true"><circle cx="12" cy="12" r="10" /><polyline points="12 6 12 12 16 14" /></svg>
          <span>{time}</span>
        </span>
        <a className="footer-docs inline-flex cursor-pointer items-center justify-self-end gap-1.25 whitespace-nowrap text-(--muted) no-underline transition-colors duration-200 hover:text-(--text)" data-wml-openURL="https://v3.wails.io" aria-label="Wails documentation">Docs
          <svg className="size-4.25" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" aria-hidden="true"><line x1="7" y1="17" x2="17" y2="7" /><polyline points="7 7 17 7 17 17" /></svg>
        </a>
      </footer>

      <div className={`toast fixed top-[max(var(--s-3),env(safe-area-inset-top))] right-[max(var(--s-3),env(safe-area-inset-right))] z-20 inline-flex max-w-[min(86vw,26.25rem)] flex-col items-start gap-0.75 rounded-[0.875rem] border border-(--glass-border) bg-[rgba(13,17,28,0.72)] px-4 py-2.75 text-[0.9375rem] leading-[1.45] text-(--text) shadow-[0_1rem_2.75rem_rgba(0,0,0,0.45)] backdrop-blur-lg pointer-events-none [transition:opacity_0.3s_ease,transform_0.35s_cubic-bezier(0.2,0.8,0.2,1)] ${isToastVisible ? 'translate-y-0 opacity-100' : '-translate-y-4 opacity-0'}`} role="status" aria-live="polite">
        <span className="text-[0.65625rem] font-bold tracking-[0.09em] text-(--muted) uppercase">From Go</span>
        <span aria-label="result" className="min-w-0">{toastMessage}</span>
      </div>
    </>
  )
}

export default GreetView

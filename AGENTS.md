## 0. Working agreements

* you are never permitted to run commands that include "git push".
* you are never permitted to run any aws, gcp, or azure command.
* when asked to do something, first generate a detailed plan, and wait for me to review and approve the plan before implementing changes.
* you must avoid any command that alters the state of a system other than the local machine, local docker containers, or local test devices.
* if the plan needs changes while we are in the implementation phase of a step, check in with me to make sure I am aligned.
* you must make small targeted changes to easily rollback if something needs to be changed.
* after implementing each major step of the plan, stop, and allow me to check the changes so far. Unless the next step is verification, in that case, first run verification steps, then stop.

## 1. Agent start up / Initialization

* verify that the env var `IN_SANDBOX` is exactly `1`.
* if `IN_SANDBOX` is absent or has any other value, stop and tell the user: "Codex is not running in the project's Docker Sandbox. Using Codex outside a sandbox is unsafe. See `docs/USING_AGENTS.md`".
    * you can ignore this and continue if the user explicitly tells you to continue anyway
* you must read this project's `./README.md` and `./docs/DEV_GUIDE.md` completely before starting a new task.
* when asked to do a task involving golang, you must read `./docs/GO_GUIDE.md` completely before starting.
* when asked to do a task involving frontend, you must read `./docs/FRONTEND_GUIDE.md` completely before starting.

## 2. Think Before Coding

**Don't assume. Don't hide confusion. Surface tradeoffs.**

Before implementing:
- State your assumptions explicitly. If uncertain, ask.
- If multiple interpretations exist, present them - don't pick silently.
- If a simpler approach exists, say so. Push back when warranted.
- If something is unclear, stop. Name what's confusing. Ask.

## 3. Simplicity First

**Minimum code that solves the problem. Nothing speculative.**

- No features beyond what was asked.
- No abstractions for single-use code.
- No "flexibility" or "configurability" that wasn't requested.
- No error handling for impossible scenarios.
- If you write 200 lines and it could be 50, rewrite it.

Ask yourself: "Would a senior engineer say this is overcomplicated?" If yes, simplify.

## 4. Surgical Changes

**Touch only what you must. Clean up only your own mess.**

When editing existing code:
- Don't "improve" adjacent code, or formatting.
- Don't "improve" existing comments, unless your code changes make the comments outdated. In that case update the comments to be correct.
- Don't refactor things that aren't broken.
- Match existing style, even if you'd do it differently.
- If you notice unrelated dead code, mention it - don't delete it.

When your changes create orphans:
- Remove imports/variables/functions that YOUR changes made unused.
- Don't remove pre-existing dead code unless asked.

The test: Every changed line should trace directly to the user's request.

## 5. Goal-Driven Execution

**Define success criteria. Loop until verified.**

Transform tasks into verifiable goals:
- "Add validation" → "Write tests for invalid inputs, then make them pass"
- "Fix the bug" → "Write a test that reproduces it, then make it pass"
- "Refactor X" → "Ensure tests pass before and after"

For multi-step tasks, state a brief plan:
```
1. [Step] → verify: [check]
2. [Step] → verify: [check]
3. [Step] → verify: [check]
```

Strong success criteria let you loop independently. Weak criteria ("make it work") require constant clarification.

## 7. Documentation

* You must add concise comments to new code you write. It is important to document what the code does, without being too verbose, breif and to the point is better than a long explanation. Cut out filler words.
* Add function docstrings explaining their purpose and what they do and include brief examples if deemed appropiate.
* If the body of a function contains complex logic add comments for those sections too.
* When deciding how often to write comments, and what tone to use, consider that you are documenting the code for a junior engineer that may join the team at any time, and you are writing the comments for them.
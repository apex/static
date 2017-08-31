/**
 * Setup.
 */

const items = document.querySelectorAll('[id]')
const links = document.querySelectorAll('.Menu a')

/**
 * Check if `el` is out out of view.
 */

function isBelowScroll(el) {
  return el.getBoundingClientRect().bottom > 0
}

/**
 * Activate item `i`.
 */

function activateItem(i) {
  links.forEach(e => e.classList.remove('active'))
  links[i].classList.add('active')
}

/**
 * Activate the correct menu item for the
 * contents in the viewport.
 */

function activate() {
  let i = 0

  for (; i < items.length; i++) {
    if (isBelowScroll(items[i])) {
      break
    }
  }

  activateItem(i)
}

/**
 * Activate scroll spy thingy.
 */

window.addEventListener('scroll', e => activate())

/**
 * Add smooth scrolling.
 */

window.addEventListener('click', e => {
  const el = e.target

  // links only
  if (el.nodeName != 'A') return

  // url
  const url = el.getAttribute('href')

  // anchors only
  if (url[0] != '#') return

  // scrolllllllll
  document.querySelector(url).scrollIntoView({
    behavior: 'smooth'
  })
})

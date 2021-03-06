// /* global */

import * as system from './system.js'
import { DB } from './db.js'

export function updateFromDB (oid, record) {
  let row = document.querySelector("div#controllers tr[data-oid='" + oid + "']")

  if (record.status === 'deleted') {
    deleted(row)
    return
  }

  if (!row) {
    row = add(oid)
  }

  const name = row.querySelector(`[data-oid="${oid}.1"]`)
  const deviceID = row.querySelector(`[data-oid="${oid}.2"]`)
  const address = row.querySelector(`[data-oid="${oid}.3"]`)
  const datetime = row.querySelector(`[data-oid="${oid}.4"]`)
  const cards = row.querySelector(`[data-oid="${oid}.5"]`)
  const events = row.querySelector(`[data-oid="${oid}.6"]`)
  const door1 = row.querySelector(`[data-oid="${oid}.7"]`)
  const door2 = row.querySelector(`[data-oid="${oid}.8"]`)
  const door3 = row.querySelector(`[data-oid="${oid}.9"]`)
  const door4 = row.querySelector(`[data-oid="${oid}.10"]`)

  row.dataset.status = record.status

  update(name, record.name)
  update(deviceID, record.deviceID)
  update(address, record.address.address, record.address.status)
  update(datetime, record.datetime.datetime, record.datetime.status)
  update(cards, record.cards.cards, record.cards.status)
  update(events, record.events.events)
  update(door1, record.doors[1])
  update(door2, record.doors[2])
  update(door3, record.doors[3])
  update(door4, record.doors[4])

  address.dataset.original = record.address.configured
  datetime.dataset.original = record.datetime.expected

  return row
}

export function onNew () {
  const records = [{ oid: '<new>', value: '' }]
  const reset = function () {}
  const cleanup = function () {}

  system.post('objects', records, reset, cleanup)
}

export function set (element, value, status) {
  const oid = element.dataset.oid
  const original = element.dataset.original
  const v = value.toString()
  const flag = document.getElementById(`F${oid}`)

  element.dataset.value = v

  if (v !== original) {
    mark('modified', element, flag)
  } else {
    unmark('modified', element, flag)
  }

  percolate(oid, modifiedX)
}

function update (element, value, status) {
  if (element && value) {
    const v = value.toString()
    const oid = element.dataset.oid
    const flag = document.getElementById(`F${oid}`)
    const previous = element.dataset.original

    element.dataset.original = v

    // check for conflicts with concurrently edited fields
    if (element.classList.contains('modified')) {
      if (previous !== v && element.dataset.value !== v) {
        mark('conflict', element, flag)
      } else if (element.dataset.value !== v) {
        unmark('conflict', element, flag)
      } else {
        unmark('conflict', element, flag)
        unmark('modified', element, flag)
      }

      percolate(oid, modifiedX)
      return
    }

    // check for conflicts with concurrently submitted fields
    if (element.classList.contains('pending')) {
      if (previous !== v && element.dataset.value !== v) {
        mark('conflict', element, flag)
      } else {
        unmark('conflict', element, flag)
      }

      return
    }

    // update fields not pending or modified
    element.value = v
    set(element, value)
  }
}

function modifiedX (oid) {
  const element = document.querySelector(`[data-oid="${oid}"]`)
  let count = 0

  if (element) {
    const list = document.querySelectorAll(`[data-oid^="${oid}."]`)
    const re = /^\.[0-9]+$/

    list.forEach(e => {
      if (e.classList.contains('modified')) {
        const oidx = e.dataset.oid
        if (oidx.startsWith(oid) && re.test(oidx.substring(oid.length))) {
          count = count + 1
        }
      }
    })

    if (count > 0) {
      element.dataset.modified = count > 1 ? 'multiple' : 'single'
      element.classList.add('modified')
    } else {
      element.dataset.modified = null
      element.classList.remove('modified')
    }
  }
}

export function rollback (row) {
  if (row && row.classList.contains('new')) {
    DB.delete('controllers', row.dataset.oid)
    system.refreshed()
  } else {
    revert(row)
  }
}

export function commit (...rows) {
  const list = []

  rows.forEach(row => {
    const oid = row.dataset.oid
    const children = row.querySelectorAll(`[data-oid^="${oid}."]`)
    children.forEach(e => {
      if (e.dataset.value !== e.dataset.original) {
        list.push(e)
      }
    })
  })

  const records = []
  list.forEach(e => {
    const oid = e.dataset.oid
    const value = e.dataset.value
    records.push({ oid: oid, value: value })
  })

  const reset = function () {
    list.forEach(e => {
      const flag = document.getElementById(`F${e.dataset.oid}`)
      unmark('pending', e, flag)
      mark('modified', e, flag)
    })
  }

  const cleanup = function () {
    list.forEach(e => {
      const flag = document.getElementById(`F${e.dataset.oid}`)
      unmark('pending', e, flag)
    })
  }

  list.forEach(e => {
    const flag = document.getElementById(`F${e.dataset.oid}`)
    mark('pending', e, flag)
    unmark('modified', e, flag)
  })

  system.post('objects', records, reset, cleanup)
}

export function add (oid) {
  const uuid = rowID(oid)
  const tbody = document.getElementById('controllers').querySelector('table tbody')

  if (tbody) {
    const template = document.querySelector('#controller')
    const row = tbody.insertRow()

    row.id = uuid
    row.classList.add('controller')
    row.classList.add('new')
    row.dataset.oid = oid
    row.dataset.status = 'unknown'
    row.innerHTML = template.innerHTML

    const commit = row.querySelector('td span.commit')
    commit.id = uuid + '_commit'
    commit.dataset.record = uuid
    commit.dataset.enabled = 'false'

    const rollback = row.querySelector('td span.rollback')
    rollback.id = uuid + '_rollback'
    rollback.dataset.record = uuid
    rollback.dataset.enabled = 'false'

    const fields = [
      { suffix: 'name', oid: `${oid}.1`, selector: 'td input.name', flag: 'td img.name' },
      { suffix: 'ID', oid: `${oid}.2`, selector: 'td input.ID', flag: 'td img.ID' },
      { suffix: 'IP', oid: `${oid}.3`, selector: 'td input.IP', flag: 'td img.IP' },
      { suffix: 'datetime', oid: `${oid}.4`, selector: 'td input.datetime', flag: 'td img.datetime' },
      { suffix: 'cards', oid: `${oid}.5`, selector: 'td input.cards', flag: 'td img.cards' },
      { suffix: 'events', oid: `${oid}.6`, selector: 'td input.events', flag: 'td img.events' },
      { suffix: 'door-1', oid: `${oid}.7`, selector: 'td select.door1', flag: 'td img.door1' },
      { suffix: 'door-2', oid: `${oid}.8`, selector: 'td select.door2', flag: 'td img.door2' },
      { suffix: 'door-3', oid: `${oid}.9`, selector: 'td select.door3', flag: 'td img.door3' },
      { suffix: 'door-4', oid: `${oid}.10`, selector: 'td select.door4', flag: 'td img.door4' }
    ]

    fields.forEach(f => {
      const field = row.querySelector(f.selector)
      const flag = row.querySelector(f.flag)

      field.id = uuid + '-' + f.suffix
      field.value = ''
      field.dataset.oid = f.oid
      field.dataset.record = uuid
      field.dataset.original = ''
      field.dataset.value = ''

      flag.id = 'F' + f.oid
    })

    return row
  }
}

function revert (row) {
  const fields = row.querySelectorAll('.field')

  fields.forEach((item) => {
    item.value = item.dataset.original
    set(item, item.dataset.original)
  })

  row.classList.remove('modified')
}

function deleted (row) {
  const tbody = document.getElementById('controllers').querySelector('table tbody')

  if (tbody && row) {
    const rows = tbody.rows

    for (let ix = 0; ix < rows.length; ix++) {
      if (rows[ix].id === row.id) {
        tbody.deleteRow(ix)
        break
      }
    }
  }
}

function mark (clazz, ...elements) {
  elements.forEach(e => {
    if (e) {
      e.classList.add(clazz)
    }
  })
}

function unmark (clazz, ...elements) {
  elements.forEach(e => {
    if (e) {
      e.classList.remove(clazz)
    }
  })
}

function percolate (oid, f) {
  let oidx = oid

  while (oidx) {
    const match = /(.*?)(?:[.][0-9]+)$/.exec(oidx)
    oidx = match ? match[1] : null
    if (oidx) {
      f(oidx)
    }
  }
}

function rowID (oid) {
  return 'R' + oid.replaceAll(/[^0-9]/g, '')
}

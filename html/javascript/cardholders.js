/* global constants */

import { getAsJSON, postAsJSON, warning, dismiss } from './uhppoted.js'

export function onEdited (event) {
  set(event.target, event.target.value)
}

export function onTick (event) {
  set(event.target, event.target.checked)

  const rowid = event.target.dataset.record
  const row = document.getElementById(rowid)

  if (row) {
    const unmodified = Array.from(row.children).every(item => !item.classList.contains('modified'))
    if (unmodified) {
      row.classList.remove('modified')
    } else {
      row.classList.add('modified')
    }
  }
}

export function onCommit (event, op) {
  if (op && op === 'add') {
    onAdd(event)
    return
  }

  onUpdate(event)
}

export function onAdd (event) {
  const id = event.target.dataset.record
  const row = document.getElementById(id)

  if (row) {
    const record = rowToRecord(id, row)

    row.classList.remove('modified')
    Array.from(row.children).forEach(item => item.classList.remove('modified'))
    busy()

    postAsJSON('/cardholders', { cardholders: [record] })
      .then(response => {
        unbusy()

        switch (response.status) {
          case 200:
            response.json().then(object => { 
              updated(object.db.updated) 
              deleted(object.db.deleted) 
            })
            break

          default:
            // reset()
            response.text().then(message => { warning(message) })
        }
      })
      .catch(function (err) {
        unbusy()
        // reset()
        warning(`Error committing update (ERR:${err.message.toLowerCase()})`)
      })
  }
}

export function onUpdate (event) {
  const id = event.target.dataset.record
  const row = document.getElementById(id)

  if (row) {
    const record = rowToRecord(id, row)

    row.classList.remove('modified')
    Array.from(row.children).forEach(item => item.classList.remove('modified'))
    busy()

    postAsJSON('/cardholders', { cardholders: [record] })
      .then(response => {
        unbusy()

        switch (response.status) {
          case 200:
            response.json().then(object => { 
              updated(object.db.updated) 
              deleted(object.db.deleted) 
            })
            break

          default:
            // reset()
            response.text().then(message => { warning(message) })
        }
      })
      .catch(function (err) {
        unbusy()
        // reset()
        warning(`Error committing update (ERR:${err.message.toLowerCase()})`)
      })
  }
}

// export function onUpdateX (event) {
//   const id = event.target.dataset.record
//   const row = document.getElementById(id)
//
//   if (row) {
//     const update = {}
//     const fields = row.querySelectorAll('.field')
//
//     fields.forEach((item) => {
//       if ((item.dataset.record === id) && (item.dataset.value !== item.dataset.original)) {
//         update[item.id] = item.dataset.value
//
//         item.parentElement.classList.add('pending')
//         item.parentElement.classList.remove('modified')
//       }
//     })
//
//     row.classList.remove('modified')
//
//     const reset = function () {
//       Object.entries(update).forEach(([k, v]) => {
//         document.getElementById(k).parentElement.classList.remove('pending')
//         document.getElementById(k).parentElement.classList.add('modified')
//       })
//
//       row.classList.add('modified')
//     }
//
//     busy()
//
//     postAsJSON('/cardholders/' + id, update)
//       .then(response => {
//         unbusy()
//
//         switch (response.status) {
//           case 200:
//             response.json().then(object => { updated(object.db.updated) })
//             break
//
//           default:
//             reset()
//             response.text().then(message => { warning(message) })
//         }
//       })
//       .catch(function (err) {
//         unbusy()
//         reset()
//         warning(`Error committing update (ERR:${err.message.toLowerCase()})`)
//       })
//   }
// }

export function onRollback (event, op) {
  if (op && op === 'delete') {
    onDelete(event)
    return
  }

  onRevert(event)
}

export function onDelete (event) {
  const id = event.target.dataset.record
  const tbody = document.getElementById('table').querySelector('table tbody')
  const row = document.getElementById(id)

  if (row) {
    const rows = tbody.rows

    for (let ix = 0; ix < rows.length; ix++) {
      if (rows[ix].id === id) {
        tbody.deleteRow(ix)
        break
      }
    }
  }
}

export function onRevert (event) {
  const id = event.target.dataset.record
  const row = document.getElementById(id)

  if (row) {
    const fields = row.querySelectorAll('.field')

    fields.forEach((item) => {
      if ((item.dataset.record === id) && (item.dataset.value !== item.dataset.original)) {
        switch (item.getAttribute('type').toLowerCase()) {
          case 'text':
          case 'number':
          case 'date':
            item.value = item.dataset.original
            break

          case 'checkbox':
            item.checked = item.dataset.original === 'true'
            break
        }
      }

      set(item, item.dataset.original)
    })

    row.classList.remove('modified')
  }
}

export function onNew (event) {
  const tbody = document.getElementById('table').querySelector('table tbody')

  if (tbody) {
    const row = tbody.insertRow()
    const name = row.insertCell()
    const controls = row.insertCell()
    const card = row.insertCell()
    const from = row.insertCell()
    const to = row.insertCell()
    const groups = []
    const uuid = 'U' + uuidv4()

    // 'constants' is a global object initialised by the Go template
    for (let i = 0; i < constants.groups.length; i++) {
      groups.push(row.insertCell())
    }

    row.id = uuid
    row.classList.add('new')

    name.style = 'border-right:0;'
    controls.style = 'border-left:0;'
    controls.classList.add('controls')

    name.innerHTML = '<img class="flag" src="images/corner.svg" />' +
                     '<input id="' + uuid + '-name" class="field name" type="text" value="" onchange="onEdited(event)" data-record="' + uuid + '" data-original="" data-value="" placeholder="-" />'

    controls.innerHTML = '<span id="' + uuid + '-commit"   class="control commit"   onclick="onCommit(event,\'add\')"   data-record="' + uuid + '" data-enabled="false">&#9745;</span>' +
                         '<span id="' + uuid + '-rollback" class="control rollback" onclick="onRollback(event,\'delete\')" data-record="' + uuid + '" data-enabled="false">&#9746;</span>'

    card.innerHTML = '<img class="flag" src="images/corner.svg" />' +
                     '<input id="' + uuid + '-card" class="field cardnumber" type="number" min="0" value="" onchange="onEdited(event)" data-record="' + uuid + '" data-original="" data-value="" placeholder="6152346" />'

    from.innerHTML = '<img class="flag" src="images/corner.svg" />' +
                     '<input id="' + uuid + '-from" class="field from" type="date" value="" onchange="onEdited(event)" data-record="' + uuid + '" data-original="" data-value="" required />'

    to.innerHTML = '<img class="flag" src="images/corner.svg" />' +
                   '<input id="' + uuid + '-to" class="field to" type="date" value="" onchange="onEdited(event)" data-record="' + uuid + '" data-original="" data-value="" required />'

    for (let i = 0; i < groups.length; i++) {
      const g = groups[i]
      const id = uuid + '-' + constants.groups[i]

      g.innerHTML = '<img class="flag" src="images/corner.svg" />' +
                    '<label class="group">' +
                    '<input id="' + id + '" class="field" type="checkbox" onclick="onTick(event)" data-record="' + uuid + '" data-original="false" data-value="false" />' +
                    '<img class="no"  src="images/times-solid.svg" />' +
                    '<img class="yes" src="images/check-solid.svg" />' +
                    '</label>'
    }
  }
}

export function onRefresh (event) {
  busy()
  dismiss()

  getAsJSON('/cardholders')
    .then(response => {
      unbusy()

      switch (response.status) {
        case 200:
          response.json().then(object => { refresh(object.db) })
          break

        default:
          response.text().then(message => { warning(message) })
      }
    })
    .catch(function (err) {
      console.log(err)
    })
}

function refresh (db) {
  updated(Object.values(db.cardholders))
}

function updated (list) {
  if (list) {
    list.forEach((record) => {
      const id = record.ID
      const row = document.getElementById(id)

      if (row) {
        row.classList.remove('new')
      }

      if (record.Name) {
        update(document.getElementById(id + '-name'), record.Name)
      }

      if (record.Card) {
        update(document.getElementById(id + '-card'), record.Card)
      }

      if (record.From) {
        update(document.getElementById(id + '-from'), record.From)
      }

      if (record.To) {
        update(document.getElementById(id + '-to'), record.To)
      }

      Object.entries(record.Groups).forEach(([k, v]) => {
        update(document.getElementById(id + '-' + k), v)
      })
    })
  }
}

function deleted (list) {
  const tbody = document.getElementById('table').querySelector('table tbody')

  if (tbody && list) {
    list.forEach((record) => {
      const id = record.ID
      const row = document.getElementById(id)

      if (row) {
        const rows = tbody.rows
        for (let i=0; i<rows.length; i++) {
          if (rows[i].id === id) {
            tbody.deleteRow(i)
            break
          }
        }
      }
    })
  }
}

function set (element, value) {
  const rowid = element.dataset.record
  const row = document.getElementById(rowid)
  const original = element.dataset.original
  const v = value.toString()
  let td = element.parentElement

  for (let i = 0; i < 10; i++) {
    if (td.tagName.toLowerCase() === 'td') {
      break
    }

    td = td.parentElement
  }

  element.dataset.value = v

  if (v !== original) {
    td.classList.add('modified')
  } else {
    td.classList.remove('modified')
  }

  if (row) {
    const unmodified = Array.from(row.children).every(item => !item.classList.contains('modified'))
    if (unmodified) {
      row.classList.remove('modified')
    } else {
      row.classList.add('modified')
    }
  }
}

function update (element, value) {
  const v = value.toString()

  if (element) {
    element.dataset.original = v

    // check for conflicts with concurrently modified fields

    if (element.parentElement.classList.contains('modified')) {
      element.parentElement.classList.remove('pending')

      if (element.dataset.value !== v.toString()) {
        element.parentElement.classList.add('conflict')
      } else {
        element.parentElement.classList.remove('modified')
        element.parentElement.classList.remove('conflict')
      }

      return
    }

    // mark fields with unexpected values after submit

    if (element.parentElement.classList.contains('pending')) {
      element.parentElement.classList.remove('pending')

      if (element.dataset.value !== v.toString()) {
        element.parentElement.classList.add('conflict')
      } else {
        element.parentElement.classList.remove('conflict')
      }
    }

    // update unmodified fields

    switch (element.getAttribute('type').toLowerCase()) {
      case 'text':
      case 'number':
      case 'date':
        element.value = v
        break

      case 'checkbox':
        element.checked = (v === 'true')
        break
    }

    set(element, value)
  }
}

function busy () {
  const windmill = document.getElementById('windmill')
  const queued = Math.max(0, (windmill.dataset.count && parseInt(windmill.dataset.count)) | 0)

  windmill.dataset.count = (queued + 1).toString()
}

function unbusy () {
  const windmill = document.getElementById('windmill')
  const queued = Math.max(0, (windmill.dataset.count && parseInt(windmill.dataset.count)) | 0)

  if (queued > 1) {
    windmill.dataset.count = (queued - 1).toString()
  } else {
    delete (windmill.dataset.count)
  }
}

function rowToRecord (id, row) {
  const name = row.querySelector('#' + id + '-name')
  const card = row.querySelector('#' + id + '-card')
  const from = row.querySelector('#' + id + '-from')
  const to = row.querySelector('#' + id + '-to')

  const record = {
    id: id,
    groups: {}
  }

  if (name && name.dataset.value !== name.dataset.original) {
    record.name = row.querySelector('#' + id + '-name').value
  }

  if (card && card.dataset.value !== card.dataset.original) {
    record.card = Number(row.querySelector('#' + id + '-card').value)
  }

  if (from && from.dataset.value !== from.dataset.original) {
    record.from = row.querySelector('#' + id + '-from').value
  }

  if (to && to.dataset.value !== to.dataset.original) {
    record.to = row.querySelector('#' + id + '-to').value
  }

  constants.groups.forEach((gid) => {
    const group = row.querySelector('#' + id + '-' + gid)
    if (group && group.dataset.value !== group.dataset.original) {
      record.groups[gid] = group.checked
    }
  })

  return record
}

// Ref. https://stackoverflow.com/questions/105034/how-to-create-a-guid-uuid
function uuidv4 () {
  return ([1e7] + -1e3 + -4e3 + -8e3 + -1e11).replace(/[018]/g, c =>
    (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
  )
}

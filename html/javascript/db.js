export const DB = {
  controllers: new Map(),

  updated: function (tag, recordset) {
    if (recordset) {
      recordset.forEach(r => update(r))
    }
  }
}

export function UpdateDB (controllers) {
  if (controllers) {
    controllers.forEach(c => {
      update(c)
    })
  }
}

function update (c) {
  const oid = c.OID

  const record = {
    OID: oid,
    name: '',
    deviceID: '',

    address: {
      address: '',
      configured: '',
      status: 'unknown'
    },

    datetime: {
      datetime: '',
      expected: '',
      status: 'unknown'
    },

    cards: {
      cards: '',
      status: 'unknown'
    },

    events: {
      events: '',
      status: 'unknown'
    },

    doors: {
      1: '',
      2: '',
      3: '',
      4: ''
    },

    status: 'unknown'
  }

  record.status = statusToString(c.Status)

  if (c.Name) {
    record.name = c.Name
  }

  if (c.DeviceID) {
    record.deviceID = c.DeviceID
  }

  if (c.IP.Address) {
    record.address.address = c.IP.Address
    record.address.configured = c.IP.Configured
    record.address.status = statusToString(c.IP.Status)
  }

  if (c.SystemTime) {
    record.datetime.datetime = c.SystemTime.DateTime
    record.datetime.expected = c.SystemTime.Expected
    record.datetime.status = c.SystemTime.Status
  }

  if (c.Cards) {
    record.cards.cards = c.Cards.Records
    record.cards.status = statusToString(c.Cards.Status)
  }

  if (c.Events) {
    record.events.events = c.Events
    record.events.status = 'ok'
  }

  if (c.Doors) {
    record.doors[1] = c.Doors[1]
    record.doors[2] = c.Doors[2]
    record.doors[3] = c.Doors[3]
    record.doors[4] = c.Doors[4]
  }

  DB.controllers.set(oid, record)
}

function statusToString (status) {
  switch (status) {
    case 1:
      return 'ok'

    case 2:
      return 'uncertain'

    case 3:
      return 'error'

    case 4:
      return 'unconfigured'
  }

  return 'unknown'
}

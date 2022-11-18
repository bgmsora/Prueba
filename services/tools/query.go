package tools

// Esta es la query para graphql de la solicitud para unidades disponibles (HasuraRequestUnitAvailable)
const QueryUnitsAvailable string = "{\"query\":\"query MyQuery {\\r\\n  mb(where: {trip_schedule_relationship: {_eq: 0}}) {\\r\\n    vehicle_id\\r\\n    position_latitude\\r\\n    position_longitude\\r\\n    trip_start_date\\r\\n    trip_id\\r\\n    position_speed\\r\\n    \\r\\n  }\\r\\n}\",\"variables\":{}}"

// Esta es la query para graphql de la solicitud por id (HasuraRequestId)
const QueryBusId string = "{\"query\":\"query MyQuery ($id:Int){\\r\\n  mb(where: {vehicle_id: {_eq: $id}}) {\\r\\n    position_latitude\\r\\n    position_longitude\\r\\n  }\\r\\n}\",\"variables\":{\"id\":XWFFF}}"

// Esta es la query para graphql de la solicitud de unidades (HasuraRequestUnits)
const QueryUnits string = "{\"query\":\"query MyQuery {\\r\\n  mb {\\r\\n    vehicle_id\\r\\n    position_latitude\\r\\n    position_longitude\\r\\n    trip_start_date\\r\\n    trip_id\\r\\n    position_speed\\r\\n    \\r\\n  }\\r\\n}\",\"variables\":{}}"

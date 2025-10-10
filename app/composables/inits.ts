export const initCustomerData = (): ICustomerData => ({
  firstName: '',
  lastName: '',
  street: '',
  housenumber: '',
  zipCode: '',
  place: '',
  email: '',
  birthdate: '',
});

export const transformTicketToCustomerData = (ticket:ITicket|TicketRecord):ICustomerData  => {
  return {
    firstName: ticket.firstName,
    lastName: ticket.lastName,
    street: ticket.street,
    housenumber: ticket.housenumber,
    zipCode: ticket.zipCode,
    place: ticket.place,
    email: ticket.email,
    birthdate: ticket.birthdate,
  };
};

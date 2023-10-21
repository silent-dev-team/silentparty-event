
export class Ticket implements ITicket {
  public firstName: string;
  public lastName: string;
  public street: string;
  public housenumber: string;
  public zipCode: string;
  public place: string;
  public email: string;
  public sold: boolean;
  public validated: boolean;
  public used: boolean;
  public filled: boolean;

  constructor(record:TicketRecord) {
    this.firstName = record.firstName;
    this.lastName = record.lastName;
    this.street = record.street;
    this.housenumber = record.housenumber;
    this.zipCode = record.zipCode;
    this.place = record.place;
    this.email = record.email;
    this.sold = record.sold;
    this.validated = record.validated;
    this.used = record.used;
    this.filled = record.filled;
  }

  get CustomerData(): ICustomerData {
    return {
      firstName: this.firstName,
      lastName: this.lastName,
      street: this.street,
      housenumber: this.housenumber,
      zipCode: this.zipCode,
      place: this.place,
      email: this.email,
    };
  }
}
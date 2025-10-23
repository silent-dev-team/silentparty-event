import { declare } from '../.nuxt/types/imports';
import { RecordModel } from 'pocketbase';
import { Ticket } from '../.nuxt/components';
declare global {
  interface ICustomerData {
    firstName: string;
    lastName: string;
    street: string;
    housenumber: string;
    zipCode: string;
    place: string;
    email: string;
    birthdate: string;
  }

  interface ITicket extends ICustomerData {
    sold: boolean;
    validated:  boolean;
    used: boolean;
    filled: boolean;
    vvk: boolean;
  }

  type TicketRecord = RecordModel & ITicket;

  interface IShopItem {
    id: string;
    title: string;
    price: number;
    description: string;
    pfand: boolean;
    pfand_item?: string;
    img?: string;
    tags?: string[];
  }

  type ShopItemRecord = RecordModel & IShopItem;

  interface IHeadPhone {
    qr: string;
    defect: boolean;
    lent: boolean;
    label: string;
  }

  type HeadPhoneRecord = RecordModel & IHeadPhone;

  interface ITicketHeadPhone {
    hp: string;
    ticket: string;
  }

  type TicketHeadPhoneRecord = RecordModel & ITicketHeadPhoneTicket;
  
  interface IAlert {
    from: string;
    msg: string;
    channel: string;
    active: boolean;
    seen: boolean;
  }

  type AlertRecord = RecordModel & IAlert;

  interface IUserStats {
    used_tickets: number;
    used_vvk: number;
    unused_vvk: number;
    used_ak: number;
    open_ak: number;
    active_ak: number;
    reserved_ak: nember;
    lent_hp: number;
    unused_hp: number;
    overbooks: number;
  }
}

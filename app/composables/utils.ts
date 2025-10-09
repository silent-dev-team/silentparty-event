export const useUtils = () => ({
  senetizeTicketCode: (code: string) => {
    const c = code.toLowerCase().match(re.ticket_number)?.[0]
    const [ ticketno, checksum ] = c?.split('-') || [ '', '' ]
    if (!ticketno || !checksum) return null
    return parseInt(ticketno).toString() + '-' + checksum.slice(0, 2)
  },
})

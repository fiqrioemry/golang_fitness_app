/* eslint-disable react/prop-types */
import { FormDelete } from "@/components/form/FormDelete";
import { useDeleteVoucherMutation } from "@/hooks/useVouchers";

const VoucherDelete = ({ voucher }) => {
  const { mutate: deleteVoucher, isPending } = useDeleteVoucherMutation();

  const handleDeleteVoucher = () => {
    deleteVoucher(voucher.id);
  };

  return (
    <FormDelete
      loading={isPending}
      title="Delete Vouchers"
      onDelete={handleDeleteVoucher}
      description="Are you sure want to delete this Voucher?"
    />
  );
};

export { VoucherDelete };

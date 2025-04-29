/* eslint-disable react/prop-types */
// src/components/address/DeleteAddress.jsx
import { FormDelete } from "@/components/form/FormDelete";
import { useDeleteAddressMutation } from "@/hooks/useUserMutation";

const DeleteAddress = ({ address }) => {
  const { mutation: deleteAddress, isLoading } = useDeleteAddressMutation();

  return (
    <FormDelete
      loading={isLoading}
      onClick={deleteAddress}
      title="Address Deletion"
      description="Are you sure want to delete this Address ?"
    />
  );
};

export default DeleteAddress;

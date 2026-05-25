import { useState, type FC } from "react";
import { Link } from "react-router";
import { useQueryClient } from "@tanstack/react-query";
import { useUsers } from "../../../hooks/user/useUsers";
import { useUserDelete } from "../../../hooks/user/useUserDelete";
import { type User } from "../../../types/user";
import DashboardLayout from "../../../layouts/DashboardLayout";
import useDocumentTitle from "../../../hooks/common/useDocumentTitle";

const UsersIndex: FC = () => {
  useDocumentTitle("Users");

  const { data: users, isLoading, isError, error } = useUsers();

  const queryClient = useQueryClient();

  const { mutate } = useUserDelete();

  const [deletingId, setDeletingId] = useState<string | null>(null);

  const handleDelete = (id: string) => {
    if (confirm("Are you sure you want to delete this user?")) {
      setDeletingId(id);
      mutate(id, {
        onSuccess: () => {
          queryClient.invalidateQueries({ queryKey: ["users"] });
        },
        onSettled: () => setDeletingId(null),
      });
    }
  };

  return (
    <DashboardLayout>
      <div className="card border-0 rounded-4 shadow-sm">
        <div className="card-header d-flex justify-content-between align-items-center">
          <span>USERS</span>
          <Link
            to="/dashboard/users/create"
            className="btn btn-sm btn-success rounded-4 shadow-sm border-0"
          >
            ADD USER
          </Link>
        </div>
        <div className="card-body">
          {/* Loading State */}
          {isLoading && (
            <div className="alert alert-info text-center">Loading...</div>
          )}

          {/* Error State */}
          {isError && (
            <div className="alert alert-danger text-center">
              Error: {error.message}
            </div>
          )}

          <div className="table-responsive">
            <table className="table table-bordered">
              <thead className="bg-dark text-white">
                <tr>
                  <th scope="col">Full Name</th>
                  <th scope="col">Email Address</th>
                  <th scope="col" style={{ width: "20%" }}>
                    Actions
                  </th>
                </tr>
              </thead>
              <tbody>
                {users?.map((user: User) => (
                  <tr key={user.id}>
                    <td>{user.name}</td>
                    <td>{user.email}</td>
                    <td className="text-center">
                      <Link
                        to={`/dashboard/users/edit/${user.id}`}
                        className="btn btn-sm btn-primary rounded-4 shadow-sm border-0 me-2"
                      >
                        EDIT
                      </Link>

                      <button
                        onClick={() => handleDelete(user.id)}
                        disabled={deletingId === user.id}
                        className="btn btn-sm btn-danger rounded-4 shadow-sm border-0"
                      >
                        {deletingId === user.id ? "DELETING..." : "DELETE"}{" "}
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </DashboardLayout>
  );
};

export default UsersIndex;

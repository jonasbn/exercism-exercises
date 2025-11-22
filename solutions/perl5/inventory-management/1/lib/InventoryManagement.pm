package InventoryManagement;

use v5.38;

sub create_inventory ($items) {
    my %inventory;

    foreach my $item (@{$items}) {
        $inventory{$item}++;
    }
    return \%inventory;
}

sub add_items ( $inventory, $items ) {
    foreach my $item (@{$items}) {
        $inventory->{$item}++;
    }    
    return $inventory;
}

sub remove_items ( $inventory, $items ) {
    foreach my $item (@{$items}) {
        next if $inventory->{$item} == 0;
        $inventory->{$item}--;
    }    
    return $inventory;
}

sub delete_item ( $inventory, $item ) {
    delete $inventory->{$item};    

    return $inventory;
}

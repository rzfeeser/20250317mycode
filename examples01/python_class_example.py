class PirateTreasure:
    def __init__(self, x=0, y=0, burried=False):
        """Initialize the treasure's coordinates."""
        self.x = x
        self.y = y
        self.burried = burried

    def set_coordinates(self, x, y):
        """Set the coordinates of the treasure."""
        self.x = x
        self.y = y
        print(f"Treasure coordinates updated to: ({self.x}, {self.y})")

    def get_coordinates(self):
        """Get the coordinates of the treasure."""
        return (self.x, self.y)

    def display_location(self):
        """Display the treasure's location."""
        print(f"Treasure is located at coordinates: ({self.x}, {self.y})")


# Example usage
if __name__ == '__main__':
    # Create a treasure instance at default coordinates
    treasure = PirateTreasure()
    treasure.display_location()

    # Set new coordinates for the treasure
    treasure.set_coordinates(10, 20)
    
    # Get and print the coordinates
    coords = treasure.get_coordinates()
    print(f"Coordinates retrieved: {coords}")

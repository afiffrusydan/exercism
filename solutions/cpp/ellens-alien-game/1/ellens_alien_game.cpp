namespace targets {

    class Alien {
        public:
            Alien(int x_init, int y_init) {
                x_coordinate = x_init;
                y_coordinate = y_init;
            }
            int get_health(){
                return health;
            }
            bool hit(){
                if (health == 0) {
                    return false;
                }
                health = health - 1;
                return true;
            }
            int x_coordinate{};
            int y_coordinate{};
            bool is_alive(){
                if (health < 1) {
                    return false;
                }
                return true;
            }
            bool teleport(int x, int y){
                x_coordinate = x;
                y_coordinate = y;
                return true;
            }
            bool collision_detection (Alien alien){
                if (x_coordinate == alien.x_coordinate && y_coordinate == alien.y_coordinate){
                    return true;
                }
                return false;
            }
        private:
            int health{3};
    };

}  // namespace targets
